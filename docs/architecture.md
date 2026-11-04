# Architecture

## Le découpage

```
cmd/evasion/     drapeaux, injection des dépendances
internal/core/    règles pures — aucune dépendance UI, réseau ni disque
internal/ai/      croyance bayésienne, IA embarquée, pilotage des bots
internal/render/  projection isométrique, contrat de formes
internal/storage/ SQLite : journal, instantanés, poids d'IA
internal/server/  WebSocket, hébergement et jonction
internal/loader/  chargement des manifestes, bac à sable WebAssembly
internal/session/ l'ordre d'une partie : qui joue, quand, ce qu'on montre
internal/text/    rendu d'une vue en caractères, saisie d'un coup
internal/preview/ aperçu SVG d'un plugin d'apparence, sans lancer le jeu
internal/schema/  génération des schémas depuis les structures Go
internal/quality/ les contrôles qui valent pour le dépôt entier
plugins/          le contenu livré, embarqué dans le binaire par go:embed
schemas/          les contrats publics, versionnés à part
```

## Quatre invariants

Ce sont les seules choses qui coûteraient cher à rétablir plus tard. Tout le
reste est révisable.

### La vue filtrée

Le noyau expose `ViewFor(side)` et l'interface ne consomme rien d'autre, **y
compris en partie locale**. Sans ça, un joueur lit la position du fugitif dans
le trafic réseau ou dans un fichier de sauvegarde, et le jeu n'a plus d'objet.

Règle de relecture : tout champ ajouté à `Game` doit être explicitement copié
dans `ViewFor`, jamais par recopie de structure. Une omission fait fuiter, un
oubli ne fait qu'afficher moins.

Un bot externe reçoit la même `View` que l'interface. Ce n'est pas une politesse :
c'est la même projection, appliquée au même endroit.

### Le déterminisme

Graine explicite portée par l'état, jamais de générateur global, IA comprise.
`core.Random` est le seul générateur autorisé, y compris pour les plugins. Il se
construit par `NewRandom(graine, flux)`, où le flux est un nom : isoler les
tirages fait qu'ajouter un dé quelque part ne déplace pas ce qui est tiré
ailleurs, et qu'une partie rejouée retrouve le même déroulement.

C'est ce qui rend possible le rejeu depuis le journal, la reproduction d'un
défaut, et la comparaison de deux versions d'IA sur les mêmes plateaux.

Piège propre à Go : le parcours d'une `map` n'a pas d'ordre stable. Tout ce qui
influence une décision passe par une tranche triée.

### La réversibilité

`Apply` a son `Undo`, effets de plugins compris. Ce n'est pas un confort
d'interface : c'est ce qui permet à l'IA d'explorer des milliers de positions
sans copier l'état à chaque nœud.

### Aucune coordonnée d'écran dans l'état

L'état ne connaît que colonne et ligne. La conversion isométrique vit dans
`internal/render`, et nulle part ailleurs.

## Le journal est la source de vérité

L'instantané n'est qu'un cache de reprise. `Resume` rejoue le journal plutôt
que de charger l'instantané : c'est ce qui vérifie en continu que le journal
reste suffisant. Un instantané chargé sans rejeu masquerait le jour où une règle
cesse d'être reproductible.

Conséquence utile : un bot non déterministe reste parfaitement jouable, puisque
le journal enregistre ses **coups** et non son état interne.

## Ce qui se recalcule ne se stocke pas

Toute information dérivable — visibilité, coups légaux, carte de croyance,
extraction en cours — est recalculée, jamais rangée dans `Game`. C'est la
condition pour que rejouer le journal reconstruise exactement le même état : un
champ dérivé qu'un `Undo` oublierait de défaire, ou qu'une reprise ne
reconstituerait pas, ferait diverger deux parties que rien ne distingue par
ailleurs.

Si un calcul devient coûteux, le cache vit **à côté** de l'état et se
reconstruit à partir de lui, jamais dedans.

## `LegalMoves` est la seule source de vérité

L'interface s'en sert pour surligner les cases, l'IA pour explorer, le serveur
pour valider ce qui arrive du réseau, le pilote de bot pour vérifier un coup
reçu. **Aucun des quatre ne réimplémente la règle.**

Un coup reçu se compare par égalité à une entrée de la liste. Le jeu ne corrige
ni n'interprète jamais un coup : un coup illégal interrompt la partie et part au
journal tel quel.

C'est aussi ce qui rend un plugin de règles utilisable par un bot qui l'ignore.
Une dépense ajoutée apparaît dans la liste, et choisir dedans suffit à jouer
correctement — un bot n'a rien à savoir des règles pour être légal.

## L'ordre de résolution est un contrat

Visibilité, contacts, capture, traces, expiration des altérations de terrain,
ressourcement, révélation, effets différés, fermeture de zone, dépense du
silence, décompte d'extraction. La liste fait foi dans `resolveTurnEnd`, et elle
ne se change pas sans changer le jeu.

Le décompte des contacts a lieu **après** le déplacement du fugitif, pas avant.
Déplacer une étape d'un cran modifie des parties entières sans qu'aucun test
unitaire ne s'en aperçoive, parce que chaque étape prise isolément reste juste.

## Le vocabulaire d'effets

Une capacité, une dépense ou un mode de jeu se décrit par composition de
primitives — `change_range`, `block_cell`, `reveal_trails` — et non par du
code.

Les cinq capacités livrées sont écrites dans ce format dès le premier jour :
c'est le test qui prouve que le vocabulaire suffit. Si l'une d'elles ne s'y
exprime pas, il manque une primitive, ce qui est une décision de contrat public
— pas un cas particulier à coder dans le noyau.

La référence complète du vocabulaire est
[`vocabulaire-effets.md`](vocabulaire-effets.md), et le format des fichiers qui
le portent est dans [`plugins.md`](plugins.md).

Un plugin ne touche jamais `*Game`. Il produit des `Effect` ou des `Move`, le
noyau les applique. C'est ce qui fait qu'il ne peut pas lire la zone scellée du
fugitif, et que `Undo` reste praticable.

**Ajouter une primitive est une décision lourde** : elle entre dans le contrat
public et n'en ressort plus sans périmer les plugins existants. Avant d'en
ajouter une, vérifier que la composition des existantes ne suffit pas.

## L'IA raisonne sur une croyance, pas sur un arbre

Information imparfaite : un minimax classique ne s'applique pas, et un élagage
alpha-bêta serait une erreur de catégorie — il n'y a pas d'état unique à
évaluer. L'approche est un filtre bayésien discret : une carte de croyance sur
les positions possibles du fugitif, propagée à chaque tour, resserrée par ce que
les inspecteurs voient et par les traces qu'ils découvrent.

**L'absence d'observation est l'information la plus riche du jeu.** Une case vue
et vide passe à zéro. C'est ce qui fait que patrouiller vaut mieux que camper —
et c'est aussi ce qu'un joueur humain doit tenir de tête pendant que son
adversaire machine l'a sous les yeux.

Les niveaux sont des jeux de poids, pas des algorithmes différents.
L'amélioration au fil des parties ajuste ces poids. Pas de réseau de neurones :
le gain serait invisible sur un jeu à quelques milliers de cases, et la carte de
croyance a l'avantage décisif de s'afficher telle quelle en vue de débogage.

L'IA livrée se branche sur la même signature qu'un bot externe. C'est ce qui
garantit que [le protocole](protocole-bot.md) suffit : si quelque chose y
manquait, le jeu ne pourrait pas jouer contre lui-même.

## Le rendu

Ebitengine. Un seul langage du noyau au pixel, pas de webview, pas de bundle
JavaScript, et les cibles web, Android et iOS depuis le même code.

Le prix, à connaître : la compilation croisée ne fonctionne que vers Windows et
WebAssembly. La cible Linux se construit en natif, sur des runners d'intégration
continue, et macOS n'est plus publiée faute de runner qui la prenne. Voir
[`construction.md`](construction.md).

Deux vues **affichées ensemble**, l'isométrique et la carte à plat à côté
d'elle. Ce n'est pas un basculement : chacune fait ce que l'autre ne sait pas
faire, et on a besoin des deux en même temps.

- **la vue isométrique** occupe l'essentiel de la fenêtre. Losanges en rapport
  2:1, tri en profondeur par `colonne + ligne`. **Elle ne montre jamais le
  plateau entier et n'essaie pas** : elle défile en suivant le jeu ;
- **la carte à plat** tient dans un panneau latéral, en permanence. Elle porte
  la vue d'ensemble, que l'isométrique abandonne : le fugitif y planifie un
  itinéraire vers une zone, les inspecteurs y répartissent une couverture. Elle
  superpose aussi les lignes de vue et les cases atteignables, illisibles en
  isométrique — les unes et les autres viennent de la vue filtrée, qui les porte
  déjà. **La carte de croyance de l'IA s'y ajoute à l'étape 9** : rien ne la
  produit avant, et l'étape 7 livre donc la carte sans elle.

**L'échelle isométrique n'est pas un intervalle mais une suite de paliers.** Le
rendu peint une planche à `RasterTileWidth` — 32 pixels par case — puis
l'agrandit vers la fenêtre par un facteur **entier** : les pixels par case valent
32, 64, 96, et rien entre les deux. Ce qui varie continûment avec la fenêtre est
le nombre de cases visibles, pas le grain.

Le facteur se calcule sur le **plancher** de surface, `MinSurfaceWidth` sur
`MinSurfaceHeight`, soit 960 sur 540 points, et non sur la surface courante. En
points et non en pixels : la fenêtre s'ouvre à la même surface visible quelle que
soit la densité de l'écran, donc l'agrandissement suit le facteur du moniteur —
un sur un écran ordinaire, deux à 150 %. Un jeu qu'on veut plus grand à
l'ouverture demande une planche plus grande, jamais une fenêtre plus grande.

**Et la garantie se simplifie au lieu de se perdre.** La vue doit contenir en
entier le champ du pion sélectionné — `Span(2*portée+1)` donne la place qu'il
demande, soit 544 sur 272 pixels de planche au plus grand préréglage. Comme la
surface vaut toujours au moins le plancher, **un champ qui tient au plancher
tient à toute taille de fenêtre** : un contrôle unique au montage remplace tout
calcul de plafond par fenêtre.

Ce contrôle porte sur l'**aire de monde** et non sur la surface : le panneau
latéral en retire sa largeur, et l'aire reçoit 632 sur 540 là où le champ en
demande 544 sur 272. Les confondre décalerait le cadrage sans qu'aucune erreur
ne le dise.

La portée qui entre dans ce calcul est celle du **préréglage**, pas celle du
tour. Une capacité qui double la vue ferait autrement sauter le cadrage à son
déclenchement, puis le rendrait au tour suivant — une caméra qui bouge seule
quand on active une capacité est déroutante, et c'est la carte à plat qui montre
ce que l'isométrique ne peut plus contenir.

**La garantie porte sur le pion sélectionné, pas sur tous.** Un inspecteur au
bord du panneau a sa ligne de vue tronquée à l'écran, et à cinq pions dispersés
sur quarante et une cases c'est le cas courant. D'où le rôle le plus important
de la carte à plat : elle porte la couverture d'ensemble du camp inspecteurs,
qui est leur information la plus stratégique et que l'isométrique ne peut pas
montrer.

Le panneau vaut 26 % de la largeur, borné entre 328 et 520 pixels — dimensionné
par ce qu'il doit montrer, soit de huit à treize pixels par case sur le plus
grand préréglage selon la largeur de la fenêtre. Les trois valeurs sont dans
`internal/render`, où un contrôle les compare à la taille du plus grand plateau.

C'est la borne basse qui commande la suite : cette couverture se rend **en teinte
de fond de case, jamais en halo autour d'un pion**, parce qu'à huit pixels par
case un halo déborderait sur les huit voisines.

**La carte à plat passe par `ViewFor` comme tout le reste.** C'est une vue de
plus, pas un accès privilégié : celle des inspecteurs ne montre pas le fugitif,
et une minicarte permanente ne donne donc aucun avantage. L'oublier ferait
fuiter la position par le chemin le plus discret du programme.

Le piège de l'isométrique est l'occlusion : les bâtiments en volume sont ce qui
rend la vue jolie et ce qui masque les cases derrière. Or le jeu porte
entièrement sur ce qu'on voit. D'où l'extrusion plafonnée à une demi-case, la
transparence des bâtiments devant un pion, et le marquage au sol des cases
visibles.
