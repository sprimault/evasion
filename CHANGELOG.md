# Journal des versions

Le format suit [Keep a Changelog](https://keepachangelog.com/fr/1.1.0/).

SemVer avec la clause du zéro : **en `0.x`, rien n'est imposé**. Le mineur
marque un jalon de la feuille de route, pas une rupture d'API — tout le reste
s'accumule en correctif, correctifs, fonctionnalités et ruptures confondus.

Quatre numéros à ne pas confondre :

| Numéro | Où | Ce qu'il suit |
|---|---|---|
| version du dépôt | tag git | le binaire |
| `shapes_version` | chaque fichier de formes | le contrat d'apparence |
| `protocol` | échanges avec un bot | le contrat de bot |
| `effects_version` | manifeste d'un plugin de règles | le vocabulaire d'effets |

Les trois derniers sont des entiers sans rapport avec SemVer. Une version peut
sortir sans qu'ils bougent ; ils ne bougent jamais sans version.

Un titre de section s'écrit `## [version] — date — titre`. La publication en
tire le nom et les notes de la version : ce qui est relu ici est ce qui sera
lu sur la page des versions, et il n'y a rien à recopier ensuite. Le titre est
facultatif ; sans lui, la version se nomme par son tag.

Chaque section est **bilingue, français d'abord, séparé par `***`** — les notes
de version sont ce que lit un auteur de plugin étranger avant de savoir s'il
doit reprendre son travail. Ce préambule reste en français : il n'est jamais
publié, et explique les conventions du dépôt à qui y contribue.

## [Non publié]

Le jeu s'appelle désormais Evasion. Aucun numéro de contrat ne bouge, donc rien
à reprendre dans un plugin ni dans un bot ; l'ancien nom de dépôt redirige.

L'accroche du dépôt et des README dit désormais ce que le fugitif cherche, là
où elle n'annonçait qu'une poursuite : le nouveau nom promet une sortie, et le
texte n'en nommait aucune.

golangci-lint et gosec embarquaient un lecteur de données d'export que Go 1.27.2
a périmé : le lint et l'analyse de sécurité échouaient sur chaque import de
chaque paquet. gosec est désormais épinglé comme golangci-lint.

***

The game is now called Evasion. No contract number changes, so nothing to
revisit in a plugin or a bot; the old repository name redirects.

The repository and README taglines now say what the fugitive is after, where
they only announced a chase: the new name promises a way out, and the text
named none.

golangci-lint and gosec embedded an export-data reader that Go 1.27.2 made
obsolete: linting and security analysis failed on every import of every package.
gosec is now pinned like golangci-lint.

## [0.6.3] — 2026-08-29 — La relecture du diff

Le contrôle qui refuse une affirmation de nombre non tenue dans les notes
reconnaît enfin les deux tournures de sa table qui se terminent par un `é`. Son
motif finissait sur une limite de mot, et l'expression régulière de Go ne la
connaît qu'en ASCII : ces deux-là n'ont donc jamais rien reconnu. Elles étaient
mortes à la naissance, et l'une a traversé huit lots sans être vue.
La table des tournures est sortie du motif, où rien ne pouvait l'atteindre, et un
test les éprouve désormais une par une.

Le chargeur tient le camp d'une capacité **passive**, ce que la 0.6.2 annonçait
sans le faire : son contrôle sortait sur la passivité avant de regarder le camp.
Une capacité passive déclarée côté fugitif se chargeait, se validait, et n'était
jamais donnée à personne — le noyau ne distribue que celles des inspecteurs. La
dispense d'une passive porte sur la phase qu'elle n'attend pas, jamais sur le
camp qu'elle sert.

Le rôle déduit du nom se borne aux cinq surcharges que le contrat nomme, au lieu
de fermer tout l'espace `inspector_`. Une forme appelée `inspector_halo` était
refusée au chargement — pas signalée, refusée — au nom d'une règle dont le
contrat promet deux fois qu'elle ne s'applique pas à un nom qu'un plugin ajoute.
C'est l'espace le plus naturel pour une surcouche par inspecteur, et il était
fermé sans que rien ne le dise.

**Le plancher de lisibilité d'un pion est de 3,81, sur un barrage, et non de 4,49
sur un lieu en recharge.** Le tableau du contrat énumérait neuf fonds quand il y
en a onze : l'ordre de peinture pose le pion par-dessus le marqueur de sa case,
et le Barreur peut fermer une case occupée. Le contrôle censé mesurer la garantie
héritait de la même omission et passait au vert. Le plancher reste au-dessus des
trois pour un que WCAG demande à un élément non textuel, mais de peu — et c'est
désormais le barrage qu'il faut regarder en premier si une palette resserre ses
couleurs. La phrase voisine, qui donne les fonds « de 17 à 230 en luminance »,
n'était juste que par accident : sa borne haute est la trace, absente du tableau.
Un contrôle la mesure à son tour.

La chute d'échelle qu'une capacité de vision doublée provoque passe de « 47 et
49 % » à « 47 puis 48 % » : elle vaut exactement `2r/(4r+1)`, et aucun préréglage
n'atteint 49. Le chiffre a été faux deux fois de suite — d'abord donné pour un
tiers, puis remesuré de travers dans le lot qui corrigeait les godoc comptant
faux. Un contrôle le tient désormais, formule contre document.

Deux affirmations de nombre de la note 0.6.2 sont rectifiées, et la version est
renommée. Le seuil des impasses valait 5, 7 et 10 pour une cible de 6, 15 et 26 :
la fraction annoncée n'en décrivait qu'un préréglage sur deux. Et le compte des
fonctions au nom français perd son quantificateur plutôt que d'être corrigé une
troisième fois — il n'y a pas de test qui le tienne, et la règle du dépôt dit
alors de ne pas l'écrire. `CONTRIBUTING` avait pris ce parti dès l'origine.

**La version s'appelle désormais « Avant le rendu ».** « Les promesses tenues »
affirmait une exhaustivité qu'une relecture de son propre diff a démentie
quelques heures après sa publication.

Le §3 de `regles.md` ne donne plus au Quartier une exclusivité que le §15 du même
document dément. Le lot qui a changé le seuil des impasses a régénéré le bloc de
mesures et corrigé la phrase jumelle du §14, sans toucher celle-ci : un plateau
de Périphérie sur mille passe lui aussi sous la cible.

***

The check that rejects unsupported claims about quantity in the notes now
recognises the two phrases in its table that end on an `é`. Its pattern ended on
a word boundary, and Go's regular expressions only know the ASCII one: those two
therefore never matched anything. They were dead on arrival, and one went through
eight batches unseen. The table of phrases has moved out of the
pattern, where nothing could reach it, and a test now exercises them one by one.

The loader now holds the side of a **passive** ability, which 0.6.2 announced
without doing: its check returned on passivity before looking at the side. A
passive ability declared on the fugitive's side loaded, validated, and was never
handed to anyone — the core only distributes the inspectors'. A passive's
exemption covers the phase it does not wait for, never the side it serves.

The role inferred from a name is now limited to the five overrides the contract
lists, instead of closing the whole `inspector_` space. A shape named
`inspector_halo` was rejected at load time — not warned about, rejected — in the
name of a rule the contract twice promises does not apply to a name a plugin
adds. It is the most natural space for a per-inspector overlay, and it was closed
without anything saying so.

**A piece's legibility floor is 3.81, on a roadblock, not 4.49 on a recharging
shelter.** The contract's table listed nine backgrounds where there are eleven:
the paint order puts the piece above its cell's marker, and the Blocker can close
an occupied cell. The check meant to measure the guarantee inherited the same
omission and passed. The floor stays above the three-to-one WCAG asks of a
non-text element, but not by much — and the roadblock is now the first thing to
look at if a palette tightens its colours. The neighbouring sentence, giving
backgrounds "from 17 to 230 in luminance", was only right by accident: its upper
bound is the trail, absent from the table. A check now measures it too.

The scale drop caused by a doubled sight ability goes from "47 and 49 %" to "47
then 48 %": it is exactly `2r/(4r+1)`, and no preset reaches 49. The figure was
wrong twice in a row — first given as a third, then measured wrong in the very
batch that fixed godoc counting wrong. A check now holds it, formula against
document.

Two claims about quantity in the 0.6.2 note are corrected, and the release is
renamed. The dead-end threshold was 5, 7 and 10 against a target of 6, 15 and 26:
"half" described only one of the two presets it named. And the count of
French-named functions loses its quantifier rather than being corrected a third
time — no test holds it, and the repository's rule then says not to write it.
`CONTRIBUTING` took that view from the start.

**The release is now called "Avant le rendu".** "Les promesses tenues" claimed a
completeness that a review of its own diff disproved hours after publication.

§3 of `regles.md` no longer grants the district an exclusivity that §15 of the
same document contradicts. The batch that changed the dead-end threshold
regenerated the measurement block and fixed the twin sentence in §14, without
touching this one: one outskirts board in a thousand also falls below the target.

## [0.6.2] — 2026-08-29 — Avant le rendu

**`protocol` passe à 4.** Un bot écrit contre la version 3 continue de jouer,
mais deux champs apparaissent dans la vue du fugitif et lui étaient jusqu'ici
indéductibles : `expense_uses`, qui donne pour chaque dépense plafonnée ce
qu'elle a consommé et ce qu'elle permet, et `decoy`, qui porte la fausse trace
armée. Les inspecteurs ne reçoivent ni l'un ni l'autre.

Les deux manquaient. `legal_moves` cessait de proposer une dépense
épuisée exactement comme il cesse d'en proposer une trop chère : mesuré, les
états « trois, deux ou un leurre restant » donnaient des vues indiscernables. Et
un leurre acheté n'est posé qu'à la résolution, le fugitif jouant encore
entre-temps — échanger sa case et sa direction laissait la vue identique octet
pour octet, si bien qu'un second achat écrasait le premier sans que rien puisse
le dire.

Les cinq champs de `core.View` restés en français passent à l'anglais dans le
même passage. Leurs noms JSON étaient déjà anglais : la sérialisation ne bouge
pas d'un octet.

Le pas gagné en repérant le fugitif n'est plus perdu quand le Coureur court sur
le même pion. Les deux touchent la mobilité, et la garde contre le double bonus
refusait l'un dès que l'autre était actif : le pas se gagnait ou non selon qu'on
jouait la capacité avant ou après le déplacement.

Les coups de capacité et de dépense nomment la case qu'ils désignent. Le Barreur
propose un barrage par voisine praticable et le leurre une trace par couple de
cases, mais la liste affichait autant de lignes identiques : le joueur choisissait
à l'aveugle.

**Le chargeur refuse une capacité ou une dépense que personne ne pourra jouer.**
Un déclenchement manquant, pris à l'autre phase, ou un camp qui ne correspond pas
à la nature de la déclaration : les trois se chargeaient sans un mot et n'étaient
jamais proposés. Le refus du déclenchement inconnu ne fermait qu'un côté du même
défaut.

`render.Grounds` retrouve l'ordre de luminance qu'il annonce : le lot qui a
reposé les couleurs de la palette a rendu le lieu actif plus clair que la zone
ouverte, corrigé l'énoncé de la palette et laissé la liste. La planche d'aperçu
s'y adosse, et sortait ses deux bandes du milieu permutées.

**Le contrat de formes tient enfin les deux gabarits qu'il promet.** Le plafond
de traits par rôle ne portait que sur l'état normal : un bâtiment à un trait
pouvait en déclarer vingt-quatre une fois surligné. Et le contrôle portait sur le
`role` quand le rendu choisit par le nom, si bien qu'une forme `building`
déclarée `piece` recevait l'emprise libre d'un pion au lieu du losange imposé —
ce que le contrat écarte nommément pour qu'un plugin ne puisse pas masquer ce que
l'adversaire doit voir. Les noms de formes que `contrat-formes.md` §6 énumère
imposent désormais leur rôle ; un nom
ajouté par un plugin garde le sien. `shapes_version` ne bouge pas : ces deux
contrôles resserrent ce qui était déjà écrit.

**`block_cell` et `open_cell` exigent désormais leur `duration`.** Le vocabulaire
promettait qu'une durée absente vaut la partie entière ; c'est vrai d'un effet
actif, faux du terrain, que la résolution purge à l'échéance — omettre le champ
donnait un tour, comme une durée de 1. Le contrat n'admet plus qu'on le taise
plutôt que de définir un défaut que personne ne devinerait.

Le contrat de formes dit le grain du sol en niveaux de luminance partout. Deux
phrases le décrivaient encore en pourcents, dont une avec une conséquence fausse
— deux couleurs séparées de douze niveaux ne se recouvrent pas, et le chargeur
les accepte. Il refuse dix niveaux et moins, ce que le document dit désormais.

La section qui définit les deux plans se contredisait sur le marqueur, qu'elle
rangeait au sol puis dans le plan vertical, et son diagramme du plan du sol
plaçait `y` vers le bas quand le schéma et le moteur le placent vers le haut. Une
flèche asymétrique écrite d'après le diagramme passait la validation et ressortait
retournée.

Les chiffres du liseré sont remesurés sur la palette livrée. Le tableau de
`contrat-formes.md` §2 datait d'avant la refonte des couleurs et n'énumérait pas
les deux lieux de
ressourcement, dont le pire fond de tous : un pion y tient à 4,49, ce que le
document donnait pour 5,92 sur une zone fermée. Les trois faces d'un bâtiment y
figurent désormais séparément, et un contrôle mesure chaque ligne.

La colonne des impasses du bloc de mesures de `regles.md` §15 se mesure contre la
cible de l'étape 4 — une par
carré de huit cases — et non contre le quart du côté du plateau, que la
génération écarte nommément. Le seuil publié valait 5, 7 et 10 quand la cible en
demande 6, 15 et 26, si bien que ses « 0,0 % » ne disaient rien. Un plateau de
Quartier sur vingt-quatre passe sous la cible, un sur mille en Périphérie.

La note de la 0.6.1 sur la règle de langue est rectifiée, dans le journal comme
sur la version publiée. Elle chiffrait l'écart et le donnait pour un défaut de
tests seulement : la proportion venait d'un comptage partiel, et des fonctions
hors test enfreignent la règle — `validerForme`, `poseUneTrace`,
`cibleHorsDuCamp` et d'autres. `CONTRIBUTING` portait la même phrase.

Un contrôle refuse désormais une affirmation de nombre non tenue dans les notes
non publiées. La règle existait pour les godoc, les règles, les README et la
palette ; le journal en était dispensé, alors qu'il devient une page publique
sans relecture.

**`THIRD-PARTY-NOTICES` porte l'union des six cibles de publication**, et non ce
que lie la machine qui l'a produit. Le jeu de modules dépend du système : sans
`GOOS` fixé, le fichier régénéré sur un poste aurait fait rougir la suite sur les
trois autres runners, et l'archive d'une plateforme serait partie sans la notice
de ce qu'elle incorpore sans les autres. Sans effet aujourd'hui, le binaire ne
liant que du Go pur, et posé pour la dépendance suivante.

Quatre godoc du noyau et du rendu comptaient faux ou nommaient ce qui n'existe
plus. Une capacité qui double la vue fait tomber le plafond d'échelle de 47 à
48 % selon le préréglage, quand la godoc annonçait un tiers — remesuré,
l'argument qu'elle servait en sort renforcé. Un sol allait
« du simple au triple » là où le
reste du dépôt dit « double et demi ». L'exclusion du noyau se justifiait par un
chiffre que les règles ont déjà corrigé. Et `Game.spend` nommait encore le
meurtre, retiré du jeu, quand `internal/text` renvoyait à une fonction que le
paquet n'a pas.

Les trois bornes du panneau de carte à plat vivent dans `internal/render` au lieu
d'un document, et un contrôle les compare à la taille du plus grand plateau. Il a
corrigé le minimum au passage : 320 pixels donnaient 7,8 pixels par case, sous
les huit dont dépend la décision de rendre la couverture en fond de case plutôt
qu'en halo. `docs/architecture.md` acte par ailleurs que la carte de croyance de
l'IA, qu'il rangeait parmi les couches de la carte, arrive à l'étape 9 — rien ne
la produit avant.

**La tête d'un inspecteur passe de cinq à six unités de rayon.** Le contrat
garantit qu'un pion garde un remplissage majoritaire ; à cinq, le noyau tombait à
48 % de la largeur du pion au plancher de rendu, sous la moitié promise. Le
contrôle qui gardait la propriété codait en dur le diamètre de la tête du
fugitif, présenté comme la plus petite forme livrée — c'était la plus grande, et
il n'exerçait donc que le cas le plus favorable. Il mesure désormais chaque pion
du contenu livré, et le tableau que le contrat publie est comparé à la mesure.

Les coefficients d'éclairage d'un prisme vivent dans `internal/render`, avec les
autres constantes du moteur. Le contrat les publie comme une règle, en disant que
c'est le genre de détail que deux implémentations règlent différemment sans qu'on
s'en aperçoive — et leur seule implémentation était une variable interne à
l'aperçu, que rien ne comparait au document : porter le dessus à 2,50 laissait la
suite verte. Un contrôle compare désormais les deux.

`internal/render` sort de l'exclusion `unused` du linter, comme le commentaire de
cette exclusion le prévoyait dès que son étape serait écrite. Aucun manquement
n'apparaît : c'était la dernière occasion où le retrait ne coûtait rien, le
moteur de rendu n'étant pas encore là pour y laisser des champs abandonnés.

***

**`protocol` moves to 4.** A bot written against version 3 keeps playing, but two
fields appear in the fugitive's view that it could not deduce: `expense_uses`,
giving for each capped expense what it has consumed and what it allows, and
`decoy`, carrying the armed false trail. Inspectors receive neither.

Both were missing. `legal_moves` stopped offering an exhausted expense exactly as
it stops offering one that costs too much: measured, the states "three, two or
one decoy left" gave indistinguishable views. And a bought decoy is only laid at
turn resolution, the fugitive still playing in between — swapping its cell and
its direction left the view identical byte for byte, so a second purchase
overwrote the first with nothing able to say so.

The five `core.View` fields still in French move to English in the same pass.
Their JSON names were already English: serialisation does not move by one byte.

The step gained by spotting the fugitive is no longer lost when the Runner is
active on the same piece. Both touch mobility, and the guard against a double
bonus refused one as soon as the other was running: the step was gained or not
depending on whether the ability was played before or after the move.

Ability and expense moves now name the cell they target. The Blocker offers one
roadblock per walkable neighbour and the decoy one trail per pair of cells, but
the list showed that many identical lines: the player was choosing blind.

**The loader now rejects an ability or expense nobody will ever be able to play.**
A missing trigger, one taken from the other phase, or a side that does not match
what is being declared: all three loaded without a word and were never offered.
Rejecting an unknown trigger only closed one side of the same defect.

`render.Grounds` is back in the luminance order it claims: the batch that reset
the palette colours made the active shelter lighter than the open zone, fixed the
palette's own wording and left the list alone. The preview board relies on it,
and was printing its two middle bands swapped.

**The shapes contract now holds both templates it promises.** The per-role stroke
cap only applied to the normal state: a building with one stroke could declare
twenty-four once highlighted. And the check keyed off `role` while rendering keys
off the name, so a `building` shape declared `piece` was given a piece's free
footprint instead of the imposed diamond — which the contract rules out by name,
so that no plugin can hide what the opponent must see. The shape names listed in
`contrat-formes.md` §6 now
impose their role; a name a plugin adds keeps its own. `shapes_version` does not
move: both checks tighten what was already written.

**`block_cell` and `open_cell` now require their `duration`.** The vocabulary
promised that a missing duration means the whole game; that holds for an active
effect, not for terrain, which turn resolution clears at its due date — omitting
the field gave one turn, same as a duration of 1. The contract no longer lets it
go unsaid, rather than defining a default nobody would guess.

The shapes contract now states the ground grain in luminance levels throughout.
Two sentences still described it as a percentage, one of them with a false
consequence — two colours twelve levels apart do not overlap, and the loader
accepts them. It rejects ten levels and below, which the document now says.

The section defining the two planes contradicted itself about markers, placing
them on the ground and then in the vertical plane, and its ground-plane diagram
put `y` downwards where the schema and the engine put it upwards. An asymmetric
arrow written from the diagram passed validation and came out flipped.

The rim figures are measured again on the shipped palette. The table in
`contrat-formes.md` §2 predated
the colour rework and did not list the two shelter grounds, one of which is the
worst background of all: a piece holds at 4.49 there, where the document gave
5.92 on a closed zone. A building's three faces now appear separately, and a
check measures every row.

The dead-end column of the measurement block in `regles.md` §15 is now measured
against the target of step 4 — one per
eight-by-eight square — rather than a quarter of the board's side, which
generation rules out by name. The published threshold was 5, 7 and 10 where the
target asks for 6, 15 and 26, so its "0.0 %" said nothing. One district board in
twenty-four falls below the target, one in a thousand on the outskirts.

The 0.6.1 note about the language rule is corrected, in the changelog and on the
published release. It put a figure on the gap and presented it as a test-only
defect: the proportion came from a partial count, and non-test functions do break
the rule — `validerForme`, `poseUneTrace`, `cibleHorsDuCamp` and others.
`CONTRIBUTING` carried the same sentence.

A check now rejects an unsupported claim about quantity in the unreleased notes.
The rule already covered godoc, the rules document, the READMEs and the palette;
the changelog was exempt, though it becomes a public page with no further
review.

**`THIRD-PARTY-NOTICES` now carries the union of the six release targets**, not
what the machine that produced it happens to link. The set of modules depends on
the system: with no `GOOS` set, a file regenerated on one machine would have made
the suite fail on the three other runners, and one platform's archive would have
shipped without the notice for what it incorporates and the others do not. No
effect today, the binary linking nothing but pure Go, and in place for the next
dependency.

Four godoc in the core and the renderer counted wrong or named what no longer
exists. An ability that doubles sight brings the scale ceiling down by nearly
half, not by a third — measured again at 47 to 48 % depending on the preset,
which strengthens the argument it served. A ground went "from single to triple"
where the rest of the repository says "double and a half". The core exclusion was
justified by a figure the rules have already corrected. And `Game.spend` still
named the murder, removed from the game, while `internal/text` referred to a
function the package does not have.

The flat map panel's three bounds now live in `internal/render` rather than in a
document, and a check compares them to the largest board's size. It corrected the
minimum along the way: 320 pixels gave 7.8 pixels per cell, below the eight the
decision to render coverage as a cell background rather than a halo depends on.
`docs/architecture.md` also records that the AI's belief map, which it listed
among the map's layers, arrives at step 9 — nothing produces it before then.

**An inspector's head goes from five to six units of radius.** The contract
guarantees that a piece keeps a majority fill; at five, the core dropped to 48 %
of the piece's width at the render floor, below the half that was promised. The
check guarding the property hard-coded the fugitive's head diameter, described as
the smallest shape shipped — it was the largest, so the check only exercised the
most favourable case. It now measures every shipped piece, and the table the
contract publishes is compared against the measurement.

A prism's face-lighting coefficients now live in `internal/render`, with the
engine's other constants. The contract publishes them as a rule, saying they are
the kind of detail two implementations settle differently without anyone
noticing — and their only implementation was a variable private to the preview,
which nothing compared to the document: setting the top face to 2.50 left the
suite green. A check now compares the two.

`internal/render` leaves the linter's `unused` exclusion, as that exclusion's own
comment planned once its step was written. Nothing is reported: this was the last
moment where the removal cost nothing, the render engine not yet being there to
leave abandoned fields behind.

## [0.6.1] — 2026-08-29 — Les textes du dépôt

**Les textes du dépôt disent ce que le code fait.** Rien ici ne change le jeu ni
les contrats : ce sont des commentaires, des godoc et des messages d'erreur qui
avaient dérivé — des noms de paquets qui en désignaient un autre, des primitives
et des numéros de contrat cités sous leurs anciens noms jusque dans les
workflows, un compte de flux aléatoires faux, et des godoc françaises où le
passage des identifiants à l'anglais avait emporté des verbes conjugués.

Deux commandes données en exemple ne marchaient pas : celle qu'un test affiche
en échouant, et le `filature valide` de la feuille de route, que le binaire
refuse depuis que la sous-commande s'appelle `validate`.

Le générateur de schémas passe enfin aux identifiants anglais, comme le reste du
dépôt, et perd un champ que rien ne lisait ni n'écrivait. Les schémas qu'il
produit sont inchangés, octet pour octet.

**`CONTRIBUTING` dit désormais la règle de langue qu'on suit, et non celle qu'on
croyait suivre.** Elle exigeait des identifiants anglais sans distinguer ce
qu'un autre paquet voit de ce qui ne quitte jamais son fichier, et personne ne la
faisait appliquer — ni dans les tests, ni dans le code livré. La contrainte porte
maintenant sur ce qui est visible depuis l'extérieur, et un helper de test suit
la langue de ce qu'il raconte.

Deux points en sortent, parce qu'ils touchaient ce qu'on lit pour décider. Les
chiffres que `plugins/base/palette.toml` commente étaient faux — le contraste
annoncé sur une zone ouverte l'était de plus de quarante pour cent —, et ils se
mesurent désormais sur la palette elle-même. Et la référence qui fige le grain
du sol remplissait sa table avec la fonction qu'elle devait garder : changer le
hachage ne la faisait pas rougir.

***

**The repository's texts say what the code does.** Nothing here changes the game
or the contracts: these are comments, godoc and error messages that had drifted —
package names pointing at a different package, primitives and contract numbers
cited under their former names down to the workflows, a wrong count of random
streams, and French godoc where the move of identifiers to English had swept up
conjugated verbs.

Two commands given as examples did not work: the one a test prints when it
fails, and the roadmap's `filature valide`, which the binary refuses now that the
subcommand is called `validate`.

The schema generator finally moves to English identifiers, like the rest of the
repository, and loses a field nothing read or wrote. The schemas it produces are
unchanged, byte for byte.

**`CONTRIBUTING` now states the language rule we follow, not the one we thought
we followed.** It demanded English identifiers without distinguishing what
another package sees from what never leaves its file, and nobody enforced it —
neither in tests nor in shipped code. The constraint now covers what is visible
from outside, and a test helper follows the language of what it tells.

Two of them stand out, because they touched what one reads to decide. The
figures `plugins/base/palette.toml` comments were wrong — the announced contrast
on an open zone by more than forty per cent — and they are now measured on the
palette itself. And the reference freezing the ground grain filled its table
with the very function it was meant to guard: changing the hash did not turn it
red.

## [0.6.0] — 2026-08-28 — Le contrat de formes

Étape 6 de la feuille de route. **Les cinq contrats publics du jeu — formes,
effets, protocole de bot, manifeste, règles — disent enfin ce que le code
fait.** Chacun annonçait quelque chose qu'il ne tenait pas : un exemple à
recopier qui ne charge pas, un déclenchement que rien ne déclenche, un refus que
seul le schéma appliquait, un chiffre qui a dérivé sans que personne le
remesure. Ce qui les gardait était la relecture ; ce sont maintenant des tests.

**`shapes_version` passe à 4 et `effects_version` à 3 : les plugins publiés sont
à reprendre.** Le meurtre quitte le vocabulaire avec sa scène de crime, les
lieux de ressourcement entrent dans la palette obligatoire, `role` devient
obligatoire dans toute forme, et trois déclenchements inertes sont retirés.

**Les règles refondues s'appliquent toutes** : capture par contact maintenu deux
résolutions, ressourcement sur quatre lieux du plateau, leurres à la place du
meurtre, cases surveillées publiques pour les deux camps, palette étagée sur la
luminance.

**Trois mesures de partie deviennent des tests.** La mobilité que le terrain
laisse à chaque camp, la part du noyau que cinq inspecteurs couvrent au meilleur
placement, et la longueur du trajet jusqu'à la zone la plus proche : trois
chiffres sur lesquels `docs/regles.md` fonde des décisions, et qu'aucun contrôle
ne tenait. Des bornes et non des valeurs, chacune renvoyant à ce qu'elle garde,
sur vingt graines en routine et deux cents derrière `-mesures`.

Le §2 y perd un chiffre périmé : cinq pions posés au mieux sans l'exclusion du
noyau en voient de six à neuf dixièmes selon le préréglage, non plus de neuf
partout. La mesure qui a produit la règle datait d'avant deux changements de
génération.

**Les exemples de plugin de la documentation se chargent.** Celui des langues
montrait un dossier `mes-traductions/` et, deux lignes plus bas, un manifeste
nommé `filature-de` — or le nom d'un plugin est celui de son dossier, et le
chargeur refuse leur désaccord. C'était le second exemple à recopier qui ne
passait pas, après celui du contrat de formes. Les trois manifestes que la
documentation publie sont désormais passés au chargeur par la suite de tests.

**Le contrat de formes annonçait un liseré deux fois trop fin.** Il le donnait
pour une unité quand la constante en vaut deux, et se contredisait vingt-cinq
lignes plus bas en écrivant « deux pixels de liseré ». Le schéma décrivait par
ailleurs des sols « du simple au triple » pour un rapport de 2,5, et la taille
de tête qui justifie le plafond d'épaisseur n'était celle d'aucun des deux
pions. Les deux derniers chiffres se mesurent désormais sur le contenu livré ;
la scène de crime, elle, quitte le dernier endroit qui la nommait encore.

**La première ligne de commande des README fonctionne.** Ils donnaient
`--preset ville` en exemple — en français jusque dans la version anglaise — et
listaient trois valeurs que le binaire refuse depuis que les clés sont passées à
l'anglais. Un test les confronte désormais aux préréglages réels. Trois autres
points corrigés au passage : l'archive contient aussi ses mentions de licence,
la planche d'aperçu montre chaque sol et non trois, et les préréglages de
plateau quittent la liste de ce qu'un plugin peut déclarer — ils sont écrits en
dur, et un manifeste qui en porterait un est refusé.

**Le vocabulaire des effets annonçait des paramètres qui n'existent pas.** Trois
primitives de zone déclaraient un champ `zone` qu'un effet ne porte pas — la
zone vient du coup —, et un manifeste qui l'écrirait est refusé pour champ
inconnu. `step` annonçait un `value` que le noyau ne consulte jamais. Les
tableaux du document sont désormais rapprochés des champs que le jeu décode :
c'était un contrat public de dix-huit primitives dont rien ne gardait les
colonnes.

**Le document qui fait foi cesse de se contredire.** Le §5 donnait le contact
pour rompu dès que le fugitif s'écarte, alors que quatre de ses huit pas le
maintiennent — et c'est un contact maintenu qui capture. Le §11 annonçait quatre
valeurs dérivées de la taille du plateau : il y en a cinq, et la période
d'étranglement manquait de son propre tableau de paramètres. Un test tient
désormais ce compte, que la godoc du noyau donnait faux depuis qu'elle avait
gagné une dérivée de plus.

**Le bloc de mesures des règles redevient ce que la mesure produit.** Il affirme
sa propre provenance — réécrit par le harnais, jamais à la main — et trois de ses
quatre colonnes ne pouvaient pas en venir : les préréglages y portaient leurs
anciens noms français, la trame agrégeait les plateaux que la validation jette,
et l'échantillon annoncé valait quatre fois celui d'une des colonnes. La trame
publiée tient désormais dans la fourchette du §3, comme la règle le promet ; la
génération, elle, n'a pas bougé. Un contrôle garde les noms — pas les chiffres,
qui sont des conséquences.

**La dernière clé française du schéma de vue passe à l'anglais.** Le contexte
d'un effet publiait `autre_pion` là où la cible qui le lit s'appelle
`other_piece` — seul nom que le passage des contrats à l'anglais avait laissé
derrière lui, dans une structure dont il avait traduit tous les autres.

**Les dictionnaires livrés rattrapent le jeu.** Le leurre et la capture
n'avaient rien à afficher, et une clé nommait encore le meurtre, retiré des
règles. La langue de repli n'a pas le droit d'avoir de trou — c'est sur elle que
retombe toute traduction incomplète —, et le contrôle qui l'alignait sur
l'anglais ne pouvait rien voir : les deux portaient le même. Le dictionnaire se
compare désormais à ce que le contenu livré déclare et à ce que le noyau
produit, dans les deux sens.

**Deux descriptions du schéma de manifeste disaient le faux.** `duration`
annonçait que zéro valait le tour courant, quand c'est l'inverse : un vaut le
tour courant et zéro vaut la partie entière, ce dont la capacité passive du
Traqueur dépend. Et le bloc de langue renvoyait à un `langue.toml` que le
chargeur ne lit jamais, le fichier s'appelant `language.toml` partout ailleurs.
Le vocabulaire des effets porte désormais la règle de durée, qui n'était écrite
que dans le schéma qui la donnait à l'envers.

**Le protocole de bot finit de passer à l'anglais.** Quatre descriptions du
schéma publié citaient encore des noms retirés — dont deux fois « bonjour » pour
un message que le même fichier nomme `hello` —, et l'exemple d'un coup montrait
`pion`, `Colonne` et `Ligne` là où le jeu envoie `piece`, `Column` et `Row`. Un
bot écrit sur cet exemple produisait un coup qui n'égalait aucune entrée de
`legal_moves`, alors que la ligne suivante exige qu'il y figure tel quel.
L'exemple de `hello` annonçait treize réglages pour dix-sept, les quatre absents
étant ceux du ressourcement et de l'étranglement.

Les exemples du document sont désormais rapprochés des structures que le jeu
sérialise, comme les énumérations du schéma l'étaient déjà.

**L'aperçu montre ce que le contrat décrit.** Il appliquait le grain du sol en
pourcents, ce que `docs/contrat-formes.md` §8 écarte nommément : sur la palette
livrée, ça suffisait à faire passer la rue sous un lieu actif et un lieu actif
sous une zone ouverte, soit les deux inversions que le refus au chargement
existe pour empêcher. Une forme d'opacité nulle s'y affichait pleine, et le
liseré s'épaississait de l'épaisseur d'un contour absent — jusqu'à six unités
au lieu de deux, son épaisseur passant sous le contrôle du plugin.

La cause commune est nommée : la constante du grain se documentait en pourcents
et se lisait en niveaux de luminance selon l'endroit. Le décalage vit désormais
dans un point unique du moteur, que le rendu et l'aperçu partagent.

**Le contrat de formes, son schéma et le chargeur disent enfin la même chose.**
Ils divergeaient à trois voix sur les variantes d'état : le schéma déclarait
`highlighted` et `out_of_sight`, le document les montrait, et le jeu décodait
une table sous une clé `variant` que ni l'un ni l'autre n'accepte. Le jeu suit
désormais le contrat publié, et une variante sous l'ancienne clé est refusée au
lieu d'être perdue en silence. Le message de refus nomme la clé telle qu'elle
s'écrit — il désignait jusqu'ici un chemin qu'aucun fichier chargeable ne
pouvait contenir.

**`role` devient obligatoire, y compris en surcharge.** C'est lui qui désigne le
gabarit, donc ce que la validation refuse : le déduire d'un nom ferait dépendre
un contrôle de jeu d'une table de noms que le contrat ne tient pas pour fermée.
L'exemple de surcharge du §4 le déclare désormais — il était refusé au
chargement pour cette raison, et pour une seconde : un de ses cercles débordait
sous le sol.

**Et `outline_thickness = 0` est refusé** au lieu d'être traité comme le défaut.
Le schéma le refusait déjà et le message d'erreur annonçait « de 1 à 4 » ; seul
le chargeur l'acceptait, faute de distinguer une valeur absente d'un zéro écrit.

**Trois déclenchements que rien ne déclenchait quittent le vocabulaire.** Fin de
tour, contact et révélation étaient ouverts aux capacités comme aux modes par le
schéma et décrits un par un dans la documentation, sans qu'aucune ligne du noyau
ne les produise : un plugin qui s'y accrochait restait inerte sans un message.
Les rétablir demandera de trancher ce qu'ils portent — lequel des pions au
contact, une révélation avant ou après la dépense du silence —, et la première
mécanique périodique retenue paiera ce coût en le sachant. Restent les deux
phases pour ce qu'un joueur déclenche, et l'étranglement pour ce que le jeu
déclenche lui-même, réservé aux modes.

**Et la phase d'une dépense compte enfin.** Les cinq dépenses livrées la
déclaraient, rien ne la lisait : une dépense annoncée sur la phase des
inspecteurs se serait proposée au fugitif.

**Le chargeur applique quatre familles de refus que le contrat publiait sans
lui.** Une durée ou un rayon négatifs, une annonce posée sur autre chose qu'un
effet différé, une échéance nulle, un mode sans nom ou sans effet, un
déclenchement inventé sur une capacité passive : tous entraient. Le schéma les
refusait déjà, ce qui est le pire des deux mondes — l'auteur qui valide son JSON
voyait un refus que le jeu ne faisait pas.

**Et deux traducteurs de la même langue se heurtent au lieu de s'écraser.** Le
code de langue était décodé puis abandonné : il n'entrait pas au registre, donc
rien ne le comparait, et celui qui chargeait en second gagnait dans l'ordre
alphabétique des dossiers.

**Rouvrir la dernière zone fermée laissait un état qui se relit autrement.** Le
retrait au milieu d'une liste ne rendait pas la tranche à nul quand elle se
vidait, alors que la troncature le faisait : `[]` au journal là où `null` était
attendu, sur un champ qui décide du rejeu octet pour octet. Le geste est
désormais écrit une seule fois.

**Un effet différé annoncé par le fugitif ne livre plus sa position.** Son
contexte porte sa case exacte, et la vue le servait tel quel aux inspecteurs :
un plugin de règles qui aurait annoncé une dépense aurait donné la position
dans le champ le plus discret du programme. Le camp caché annonce désormais
qu'un effet vient et quand, jamais où — les inspecteurs, dont les positions
sont publiques, gardent le leur.

**Le silence se voit, se compte et se perd.** Trois défauts du même canal
d'information. Les inspecteurs n'apprenaient jamais qu'un silence avait été
payé, le drapeau retombant avant qu'ils reprennent la main : la vue porte
désormais le tour de l'achat. Le compte à rebours annonçait une période entière
au tour même de la révélation, c'est-à-dire à l'instant où le fugitif décide de
payer. Et un silence sans emploi se reportait sans borne, ce qui en faisait une
assurance qu'il valait toujours mieux prendre tôt — il couvre le tour de son
achat, et se perd s'il n'a rien couvert.

**Naître sur un lieu de ressourcement n'est pas y entrer.** Le fugitif dont le
tirage de départ tombait sur un lieu y gagnait deux points à la première
résolution sans avoir bougé, et consommait le lieu pour huit tours. La règle
attache le gain à l'entrée : il n'a rien payé pour être là, donc le lieu reste
actif jusqu'à ce qu'il en reparte et y revienne.

**Le Barreur ferme la case que le joueur choisit.** Sa capacité n'en portait
aucune, et le barrage tombait invariablement dans le coin du plateau, sur la
zone d'extraction qui s'y trouve. Il pose désormais son obstacle sur l'une des
huit cases voisines de son pion : barrer à distance retirerait une case d'entrée
d'une zone sans dégarnir la couverture, que la règle fait payer par un pion
posté dessus.

**Les sols et les pions se distinguent par leur luminance, et le chargeur le
vérifie.** Trois couleurs se tenaient en trois niveaux de gris — une zone
ouverte, un lieu actif et le fugitif — et deux autres en neuf. Les dix sont
reposées d'un coup, étagées de dix-sept niveaux au moins, trente autour du
fugitif : un sol confondu avec un autre se rattrape sur la carte à plat, un
camp confondu avec l'autre fausse la lecture du plateau.

Le grain du sol déplaçant chaque case de cinq niveaux, deux sols voisins qui
s'en approchent échangeraient leur rang à l'affichage. Toute palette qui les
resserre sous dix est désormais refusée au chargement, la paire fautive nommée.

**Le fugitif voit ce que les inspecteurs surveillent.** Les cases tenues par une
ligne de vue ne partaient qu'à un camp, alors que les positions des pions et le
terrain sont publics : ce calcul se déduit, mais le dérouler de tête à chaque
coup adverse ne handicape que celui des deux joueurs qui fatigue.

**Le Guetteur double sa portée au lieu de la tripler.** Il déclarait huit cases
en absolu, juste tant qu'un seul préréglage existait : sur un Quartier, la portée
vaut quatre et la capacité la portait à douze. Les effets acceptent désormais un
`mode`, absent pour additionner comme avant, `multiply` pour multiplier — une
valeur absolue ne peut plus dire une intention relative depuis que la portée
dérive de la taille du plateau.

**Deux refus de plus au chargement, et une promesse retirée.** Une cible qui ne
désigne personne — `other_piece` ou `all_pieces` dans une dépense du fugitif,
qui est seul — et un code de langue hors de l'étiquette BCP 47 étaient tous deux
acceptés. En revanche, la promesse qu'une cible « incompatible avec le camp
déclarant » serait refusée disparaît : viser le camp adverse est le cas
ordinaire, et distinguer ce qui avantage le fugitif de ce qui le gêne
demanderait au chargeur de juger l'intention de chaque effet. Un plugin qui
déclare côté inspecteurs un effet lui rendant de la résistance se charge : il
est mal écrit, pas invalide.

**Le Chef ordonne un coup de filet.** Sa capacité partageait la vue d'un
coéquipier, ce qui ne voulait rien dire : les cinq inspecteurs sont un joueur
unique, et la vue du camp unit déjà ce que chacun voit. Elle rend désormais la
position du fugitif publique deux tours plus tard, annoncée aux deux camps —
seule capacité d'information à côté de quatre capacités de terrain, elle donne
aux inspecteurs le moyen d'en acheter que le fugitif avait déjà trois fois pour
la brouiller.

**Un silence couvre le tour, pas une révélation.** Un coup de filet peut tomber
le même tour que la révélation périodique, et le fugitif paie avant que les
inspecteurs jouent : il ne peut pas anticiper la coïncidence, donc il n'a pas à
la payer deux fois. Il s'achète contre être trouvé, jamais contre se montrer —
une révélation qu'il s'inflige lui-même reste en vigueur.

**Le meurtre disparaît, et la scène de crime avec lui.** Le leurre fait le même
travail en mentant au lieu de se dénoncer, ce qui laisse au fugitif quelque
chose à faire de son information au lieu de la sacrifier. Le jeu quitte le
registre du polar noir, et c'est le prix. Un plugin de règles peut le rétablir —
le coût et le plafond sont déclaratifs — mais il devra livrer la forme et la
couleur de sa scène. `shapes_version` passe à 4 pour cette raison, et
`mark_crime_scene` quitte le vocabulaire.

**Défaire une dépense rendait un état équivalent, pas identique.** Le compteur
d'emplois gardait sa clé à zéro, et la table restait allouée là où elle était
nulle : sans effet sur les règles, puisqu'une clé absente se lit zéro, mais
l'IA compare des états pour reconnaître une position déjà explorée et en aurait
vu deux. Le geste qui manquait était écrit deux fois ailleurs dans le noyau ; il
est désormais écrit une seule.

**Le fugitif peut mentir avec ses traces.** Un leurre coûte un point, trois fois
par partie, et substitue une fausse trace à toutes celles que son déplacement
aurait laissées ce tour-là — il ne s'y ajoute pas, faute de quoi un inspecteur
qui en découvrirait deux incompatibles saurait qu'il y a eu leurre. Effacer
laisse un trou, et un trou est déjà une information ; une fausse trace attaque
la déduction elle-même.

Le joueur choisit la case et la direction parmi ce que le terrain autorise : un
leurre a la forme d'un pas qui n'a pas eu lieu, donc il ne peut produire qu'une
trace qu'un vrai déplacement aurait pu laisser. `effects_version` passe à 3.

**Le chargeur applique trois refus qu'il annonçait sans les faire.** Un nom de
plugin hors du motif publié, une licence absente de la liste fermée, un
déclenchement inventé ou `strangling` sur une capacité étaient tous acceptés.
La licence est le cas le plus net : la liste blanche existe pour écarter les
mentions qu'on ne peut pas juger, et « à voir » entrait sans un mot.

**Les schémas de `schemas/` sont désormais exécutés.** Un test valide le
manifeste livré contre le contrat publié, et une série de manifestes ne portant
chacun qu'une faute, pour que chaque refus soit exercé séparément. Tant que
personne ne les exécutait, ils pouvaient mentir comme un commentaire.

**Les préréglages s'appellent `district`, `outskirts` et `city`.** Leurs clés
étaient en français quand les dictionnaires livrés déclaraient déjà les libellés
sous les noms anglais : aucun préréglage n'aurait trouvé son libellé, et le
repli sur le français n'aurait pas joué non plus, la clé manquant des deux
côtés. `filature -preset quartier` devient `filature -preset district`.

**Le protocole de bot passe à la version 3, tout en anglais.** Ses six types de
message étaient en français sauf un, que le document et le schéma nommaient
d'ailleurs différemment : un auteur de bot ne pouvait deviner ni la règle ni
l'exception. Ils sont désormais `hello`, `ready`, `play`, `move`, `over` et
`error`, comme tous les identifiants publics du projet.

Le refus d'un bot qui annoncerait une autre version est écrit et exercé, mais il
n'entrera en jeu qu'avec le pilote de bots, à l'étape 9 : rien ne l'appelle
encore, et la suite de tests le compte comme la seule mécanique inerte du
dépôt.

**Les trois numéros de contrat concordent à nouveau avec ce que le code
applique.** Le schéma de manifeste figeait `effects_version` à 1 quand le noyau
appliquait la 2, et celui du protocole restait à 1 dans ses deux messages. Le
même schéma annonçait quatre motifs de fin francisés là où le jeu en produit
cinq en anglais, et sa clause conditionnelle sur les effets différés testait un
nom de primitive qui n'existe pas — elle refusait donc la seule forme valide de
celle qu'elle prétendait contraindre.

**Les cinq sols entrent tous dans le contrat d'apparence.** Les deux couleurs
des lieux de ressourcement étaient livrées dans la palette sans être exigées
d'un plugin : une palette tierce qui les omettait était acceptée, et le rendu
cherchait un nom absent sur toutes les cases d'abri. La planche de `filature
preview` les montre désormais — elle n'en affichait que trois, et laissait donc
de côté les deux fonds que la palette livrée signale elle-même comme trop
proches l'un de l'autre.

**Les zones se ferment aux tours annoncés.** Le préavis de deux tours était
déclaré par le mode d'étranglement du contenu livré, et s'ajoutait à une cadence
que le noyau calculait déjà en tours de fermeture : tout glissait de deux tours,
et la dernière zone tombait au dernier tour de la partie au lieu de laisser de
quoi en profiter. Le préavis devient un paramètre de partie, refusé s'il dépasse
la période — voir une fermeture venir est une règle, et une règle ne se négocie
pas par manifeste.

**Le taux de rues ne juge plus que la trame.** Les blocs percés par-dessus —
zones d'extraction et lieux de ressourcement — n'entrent plus dans la mesure :
ce ne sont pas du terrain mais des dispositifs, publics et désignés, et les
agréger faisait dépendre le plafond d'une composition qui varie avec la taille
du plateau. Le Quartier passait ainsi la borne haute par ses seules zones, si
bien que la génération y jetait quatre-vingt-douze tirages sur cent et ne
retenait que les trames les plus denses. Il en sort plus ouvert qu'il ne
l'était, un tirage suffit là où il en fallait quatorze, et quelques-uns de ses
plateaux portent moins de cinq impasses.

**Cinq règles ne s'appliquaient pas.** L'étranglement ne fermait aucune zone, un
barrage et un percement duraient jusqu'à la fin de la partie, les capacités
passives n'étaient jamais posées, le déplacement rendu par un repérage
n'existait pas, et un fugitif immobile sur un lieu de ressourcement s'y
rechargeait chaque fois que celui-ci redevenait actif.

**La résistance se récupère, sur quatre lieux du plateau.** Un lieu de
ressourcement est un bloc au même format qu'une zone d'extraction, posé dans la
couronne entre le noyau de départ et les sorties. Le fugitif y regagne deux
points en fin de tour, et le lieu entier passe en recharge pour huit tours — les
deux camps voyant lequel a servi et quand il revient. Il paie donc ses points en
information, comme pour ses autres dépenses.

**La trame de rues ne se resserre plus sur le dernier îlot.** Le quatrième bord
était ajouté comme avenue sans regarder la distance au dernier axe tiré.

**La capture entre dans le jeu.** Un inspecteur qui reste au contact du fugitif
deux résolutions de suite le rattrape. En terrain ouvert celui-ci rompt le
contact d'une diagonale ; dans un couloir d'une case il ne le peut pas, et le
double déplacement redevient une dépense de survie. Un inspecteur qui occupe sa
case compte désormais comme au contact — l'en empêcher lui aurait appris où il
se trouve.

**Une zone ne se neutralise plus, ses cases s'occupent.** Un pion posté dessus
n'en retire qu'une case d'entrée : fermer une zone demande autant de pions
qu'elle a de cases de rue. L'étranglement s'arrête en laissant trois zones
ouvertes, et sa cadence suit la durée de la partie au lieu d'être figée.

**Quand deux conditions de fin tombent ensemble, celle du fugitif l'emporte.**
Les inspecteurs en ont quatre, il n'en a qu'une.

**Le noyau de départ suit la taille du plateau, et les inspecteurs ne s'y
placent plus.** Son rayon valait cinq quelle que soit la ville, pendant que la
portée de vue, elle, suivait le côté : cinq pions posés autour du centre y
voyaient le fugitif dès le premier tour dans la quasi-totalité des parties, et
d'autant mieux que le plateau était grand. `Settings` porte donc un
`centre_radius`, que les bots reçoivent dans `hello`.

**Les plugins d'apparence se chargent et se valident.** Un plugin déclare ce
qu'il remplace, le reste retombe sur le contenu livré. Une forme qui déborde de
son gabarit, un polygone hors de ses trois à trente-deux sommets, une couleur
absente de la palette, une valeur hexadécimale écrite dans une forme, un prisme
hors du rôle bâtiment : chacun arrête le chargement en nommant le chemin
complet de la clé fautive, et tous sont listés d'un coup.

Le contrôle vaut **au démarrage**, plugin local compris, et pas seulement pour
`filature validate` — une forme qui déborde masque les cases voisines, ce qui
est un avantage de jeu déguisé en habillage.

Deux plugins qui redéfinissent la même forme sont un conflit, nommé des deux
côtés. Surcharger le contenu livré n'en est pas un.

**`filature preview <dossier>` rend un plugin d'apparence en SVG**, sans lancer
le jeu : la planche de ses formes sur chaque sol, et un plateau en
situation. Le plugin est fusionné sur le contenu livré avant d'être rendu, et la
planche marque ce qui vient de lui — une clé mal orthographiée passe la
validation sans rien surcharger, et rien d'autre ne le dirait.

**L'archive `js/wasm` n'est plus publiée.** Elle ne contenait que le `.wasm`,
inutilisable sans `wasm_exec.js` ni page HTML. La cible continue d'être
compilée : c'est le seul contrôle qui empêche une dépendance d'introduire du
cgo sans que rien ne le signale.

**Les archives portent `THIRD-PARTY-NOTICES`**, à côté de `LICENSE` et
`NOTICE`. Les licences des bibliothèques liées au binaire l'exigent à la
redistribution, et les versions précédentes ne le faisaient pas. Le fichier est
généré depuis les dépendances réelles ; un test échoue quand il vieillit.

***

Step 6 of the roadmap. **The game's five public contracts — shapes, effects, bot
protocol, manifest, rules — finally say what the code does.** Each announced
something it did not hold: a copy-paste example that would not load, a trigger
nothing triggers, a rejection only the schema applied, a figure that drifted
without anyone measuring it again. What guarded them was proofreading; they are
now guarded by tests.

**`shapes_version` moves to 4 and `effects_version` to 3: published plugins need
revisiting.** Murder leaves the vocabulary along with its crime scene, respite
places enter the mandatory palette, `role` becomes mandatory on every shape, and
three inert triggers are dropped.

**The reworked rules all apply**: capture by contact held across two
resolutions, respite on four places on the board, decoys in place of murder,
watched cells public to both sides, palette tiered by luminance.

**Three gameplay measurements become tests.** The mobility the terrain leaves
each side, the share of the starting core five inspectors cover at their best
placement, and the distance to the nearest extraction zone: three figures
`docs/regles.md` bases decisions on, and that no check held. Bounds rather than
values, each pointing at what it guards, over twenty seeds by default and two
hundred behind `-mesures`.

§2 loses a stale figure along the way: five pieces placed at their best without
the core exclusion see six to nine tenths of it depending on the preset, not more
than nine everywhere. The measurement that produced the rule predated two
generation changes.

**The documentation's plugin examples load.** The language one showed a
`mes-traductions/` folder and, two lines below, a manifest named `filature-de` —
but a plugin's name is its folder's, and the loader refuses their disagreement.
It was the second copy-paste example that would not pass, after the shape
contract's. All three manifests the documentation publishes now go through the
loader in the test suite.

**The shape contract announced a rim twice too thin.** It gave one unit where
the constant is two, contradicting itself twenty-five lines later with "two
pixels of rim". The schema also described grounds ranging "one to three" for a
ratio of 2.5, and the head size justifying the thickness cap belonged to neither
piece. The last two figures are now measured on the shipped content; the crime
scene leaves the final place that still named it.

**The READMEs' first command line works.** They gave `--preset ville` as an
example — in French even in the English version — and listed three values the
binary refuses now that the keys are English. A test checks them against the
real presets. Three more points fixed along the way: the archive also holds its
licence notices, the preview sheet shows every ground rather than three, and
board presets leave the list of what a plugin can declare — they are hard-coded,
and a manifest carrying one is refused.

**The effects vocabulary announced parameters that do not exist.** Three zone
primitives declared a `zone` field an effect does not carry — the zone comes
from the move — and a manifest writing it is refused for an unknown field.
`step` announced a `value` the core never reads. The document's tables are now
checked against the fields the game decodes: this was a public contract of
eighteen primitives whose columns nothing guarded.

**The authoritative document stops contradicting itself.** §5 held contact
broken as soon as the fugitive steps away, when four of his eight steps keep it
— and it is sustained contact that captures. §11 announced four values derived
from board size: there are five, and the strangling period was missing from its
own parameter table. A test now keeps that count, which the core's godoc had
wrong ever since it gained one more derived value.

**The rules' measurement block becomes what the measurement produces again.** It
asserts its own provenance — rewritten by the harness, never by hand — and three
of its four columns could not have come from it: the presets carried their old
French names, the street ratio aggregated boards validation throws away, and the
announced sample was four times one column's own. The published ratio now sits
inside the §3 range, as the rule promises; generation itself has not changed. A
check guards the names — not the figures, which are consequences.

**The last French key in the view schema turns English.** An effect's context
published `autre_pion` where the target reading it is called `other_piece` — the
only name the contracts' move to English had left behind, in a structure whose
every other field it had translated.

**The shipped dictionaries catch up with the game.** The decoy and the capture
had nothing to display, and one key still named murder, dropped from the rules.
The fallback language is not allowed to have holes — every incomplete
translation falls back on it — and the check aligning it with English could see
nothing: both had the same hole. The dictionary is now compared against what the
shipped content declares and what the core produces, both ways.

**Two manifest schema descriptions said the opposite of the truth.** `duration`
announced that zero meant the current turn, when it is the other way round: one
means the current turn and zero means the whole game, which the Tracker's
passive ability relies on. And the language block pointed at a `langue.toml` the
loader never reads, the file being called `language.toml` everywhere else. The
effects vocabulary now carries the duration rule, which was written only in the
schema that had it backwards.

**The bot protocol finishes its move to English.** Four descriptions in the
published schema still cited retired names — twice "bonjour" for a message the
same file calls `hello` — and the move example showed `pion`, `Colonne` and
`Ligne` where the game sends `piece`, `Column` and `Row`. A bot written from
that example produced a move matching no entry in `legal_moves`, though the very
next line requires it to appear there as is. The `hello` example announced
thirteen settings out of seventeen, the four missing ones belonging to respites
and strangling.

The document's examples are now checked against the structures the game
serialises, as the schema's enumerations already were.

**The preview shows what the contract describes.** It applied the ground grain
as a percentage, which `docs/contrat-formes.md` §8 rules out by name: on the
shipped palette that was enough to push the street below an active respite and
an active respite below an open zone — the very two rank swaps the load-time
refusal exists to prevent. A stroke with zero opacity drew solid, and the rim
thickened by the width of an absent outline, up to six units instead of two,
putting its width under the plugin's control.

The common cause is named: the grain constant documented itself in percent and
was read in luminance levels depending on the caller. The shift now lives in one
place in the engine, shared by the renderer and the preview.

**The shape contract, its schema and the loader now say the same thing.** They
diverged three ways on state variants: the schema declared `highlighted` and
`out_of_sight`, the document showed them, and the game decoded a table under a
`variant` key neither accepts. The game now follows the published contract, and
a variant under the old key is refused instead of silently lost. The rejection
message names the key as it is written — until now it pointed at a path no
loadable file could contain.

**`role` becomes mandatory, overrides included.** It designates the template,
hence what validation refuses: deriving it from a name would make a gameplay
check depend on a name table the contract does not hold closed. The §4 override
example now declares it — it was refused at load time for that reason, and for a
second one: one of its circles dipped below the ground.

**And `outline_thickness = 0` is refused** instead of being treated as the
default. The schema already refused it and the error message announced "1 to 4";
only the loader accepted it, unable to tell an absent value from a written zero.

**Three triggers that nothing triggered leave the vocabulary.** Turn end,
contact and reveal were open to both abilities and modes in the schema and
described one by one in the documentation, without a single line of the core
producing them: a plugin hooking onto them stayed inert without a message.
Restoring them will mean deciding what they carry — which of the pieces in
contact, a reveal before or after the silence is spent — and the first periodic
mechanic kept will pay that cost knowingly. What remains: the two phases for
what a player triggers, and strangling for what the game triggers itself,
reserved to modes.

**And an expense's phase finally counts.** All five shipped expenses declared
it, nothing read it: an expense announced on the inspectors' phase would have
been offered to the fugitive.

**The loader applies four families of rejection the contract published without
it.** A negative duration or radius, an announcement on something other than a
deferred effect, a zero deadline, a mode with no name or no effect, an invented
trigger on a passive ability: all went through. The schema already refused them,
which is the worst of both worlds — an author validating their JSON saw a
rejection the game did not make.

**And two translators of the same language now collide instead of overwriting.**
The language code was decoded then dropped: it never reached the registry, so
nothing compared it, and whoever loaded second won by folder order.

**Reopening the last closed zone left a state that reads back differently.**
Removing from the middle of a list did not return the slice to nil when it
emptied, where truncation did: `[]` in the journal where `null` was expected, on
a field that decides byte-for-byte replay. The gesture is now written once.

**A deferred effect announced by the fugitive no longer gives away his
position.** Its context carries his exact cell, and the view served it to the
inspectors as is: a rules plugin announcing a fugitive expense would have handed
over the position through the quietest field in the program. The hidden side now
announces that an effect is coming and when, never where — inspectors, whose
positions are public, keep theirs.

**Silence is seen, counted and lost.** Three faults on the same information
channel. Inspectors never learnt that a silence had been paid, the flag dropping
before they took their turn: the view now carries the turn it was bought on. The
countdown announced a full period on the very turn of the reveal, that is, at
the moment the fugitive decides whether to pay. And an unused silence carried
over indefinitely, making it insurance always worth buying early — it now covers
the turn it is bought on, and is lost if it covered nothing.

**Being born on a respite is not entering it.** A fugitive whose starting draw
landed on a respite gained two points at the first resolution without having
moved, and spent the place for eight turns. The rule ties the gain to entering:
he paid nothing to be there, so the respite stays active until he leaves and
comes back.

**The Blocker closes the cell the player picks.** Its ability carried none, so
the roadblock always landed in the board's corner, on the extraction zone that
sits there. It now drops the obstacle on one of the eight cells around its
piece: blocking at range would take an entry cell away from a zone without
thinning the cover, which the rules make you pay for with a piece posted on it.

**Grounds and pieces are told apart by luminance, and the loader checks it.**
Three colours sat within three grey levels of each other — an open zone, an
active respite and the fugitive — and two more within nine. All ten are reset at
once, spaced at least seventeen levels apart and thirty around the fugitive: one
ground mistaken for another is caught on the flat map, one side mistaken for the
other misreads the whole board.

Since the ground grain shifts each cell by five levels, two neighbouring grounds
closer than that would swap ranks on screen. Any palette packing them under ten
is now refused at load time, with the offending pair named.

**The fugitive now sees what the inspectors watch.** Cells held by a line of
sight went to one side only, though piece positions and terrain are public: the
calculation follows from them, but running it in your head on every opposing
move only handicaps whichever of the two players gets tired.

**The Lookout doubles his range instead of tripling it.** He declared eight cells
outright, correct while a single preset existed: on a Small board range is four
and the ability took it to twelve. Effects now accept a `mode` — absent to add as
before, `multiply` to multiply — since an absolute value can no longer express a
relative intent once range derives from board size.

**Two more rejections at load time, and one promise withdrawn.** A target that
designates nobody — `other_piece` or `all_pieces` in a fugitive expense, he being
alone — and a language code outside the BCP 47 tag were both accepted. The
promise that a target "incompatible with the declaring side" would be refused, on
the other hand, is gone: targeting the opposing side is the ordinary case, and
telling what helps the fugitive from what hinders him would require the loader to
judge each effect's intent. A plugin declaring an inspector-side effect that
restores his stamina loads: it is badly written, not invalid.

**The Chief calls a dragnet.** His ability shared a teammate's view, which meant
nothing: the five inspectors are a single player, and the side's view already
merges what each of them sees. It now makes the fugitive's position public two
turns later, announced to both sides — the only information ability next to four
terrain ones, giving the inspectors a way to buy what the fugitive already had
three ways to blur.

**Silence covers the turn, not a reveal.** A dragnet can land on the same turn as
the periodic reveal, and the fugitive pays before the inspectors play: he cannot
anticipate the coincidence, so he does not pay for it twice. It is bought against
being found, never against showing oneself — a reveal he inflicts on himself
still stands.

**Murder is gone, and the crime scene with it.** The decoy does the same work by
lying instead of giving itself away, which leaves the fugitive something to do
with his information rather than sacrificing it. The game leaves the noir
register, and that is the price. A rules plugin can bring it back — cost and cap
are declarative — but it will have to ship the shape and the colour of its
scene. `shapes_version` moves to 4 for that reason, and `mark_crime_scene`
leaves the vocabulary.

**Undoing an expense returned an equivalent state, not an identical one.** The
use counter kept its key at zero, and the table stayed allocated where it had
been nil: no effect on the rules, since a missing key reads as zero, but the AI
compares states to recognise an already-explored position and would have seen
two. The missing gesture was written twice elsewhere in the core; it is now
written once.

**The fugitive can now lie with his trails.** A decoy costs one point, three
times a game, and substitutes a false trail for every one his movement would
have left that turn — it does not add to them, otherwise an inspector finding
two incompatible ones would know a decoy had been played. Erasing leaves a gap,
and a gap is already information; a false trail attacks the deduction itself.

The player picks the cell and the direction from what the terrain allows: a
decoy has the shape of a step that never happened, so it can only produce a
trail a real move could have left. `effects_version` moves to 3.

**The loader now applies three rejections it announced without performing.** A
plugin name outside the published pattern, a licence missing from the closed
list, an invented trigger or `strangling` on an ability were all accepted. The
licence is the clearest case: the allow-list exists to keep out entries nobody
can judge, and "à voir" walked straight in.

**The schemas in `schemas/` are now executed.** A test validates the shipped
manifest against the published contract, along with a series of manifests each
carrying a single fault, so every rejection is exercised on its own. As long as
nothing ran them, they could lie like a comment.

**The presets are now called `district`, `outskirts` and `city`.** Their keys
were in French while the shipped dictionaries already declared the labels under
the English names: no preset would have found its label, and the fallback to
French would not have helped either, the key being missing on both sides.
`filature -preset quartier` becomes `filature -preset district`.

**The bot protocol moves to version 3, entirely in English.** Its six message
types were in French except one, which the document and the schema happened to
name differently: a bot author could guess neither the rule nor the exception.
They are now `hello`, `ready`, `play`, `move`, `over` and `error`, like every
other public identifier in the project.

Turning away a bot that announced a different version is written and exercised,
but it only comes into play with the bot driver, at step 9: nothing calls it
yet, and the test suite counts it as the repository's only inert mechanism.

**The three contract numbers agree again with what the code applies.** The
manifest schema pinned `effects_version` to 1 while the core applied 2, and the
protocol schema stayed at 1 in both its messages. That same schema announced
four French end reasons where the game produces five English ones, and its
conditional clause on deferred effects tested a primitive name that does not
exist — so it rejected the only valid form of the very primitive it claimed to
constrain.

**All five grounds are now part of the appearance contract.** The two respite
colours shipped in the palette without being required of a plugin: a
third-party palette that omitted them was accepted, and rendering then looked up
a missing name on every respite cell. The `filature preview` board now shows
them — it displayed only three, and so left out the very two backgrounds the
shipped palette itself flags as too close to each other.

**Zones now close on the turns announced.** The two-turn notice was declared by
the shipped strangling mode, and added itself to a cadence the core already
computed as closing turns: everything slid by two, and the last zone fell on the
final turn of the game instead of leaving room to use it. The notice becomes a
game setting, refused if it exceeds the period — seeing a closure coming is a
rule, and a rule is not negotiable by manifest.

**The street ratio now judges the grid alone.** The blocks carved on top —
extraction zones and respites — no longer count towards it: they are not
terrain but fixtures, public and marked, and lumping them together made the
ceiling depend on a composition that varies with board size. The Small preset
was therefore over the upper bound on its zones alone, so generation threw away
ninety-two draws out of every hundred and kept only the densest grids. It comes
out more open than it was, one draw now suffices where fourteen were needed,
and a few of its boards carry fewer than five dead ends.

**Five rules did not apply.** Strangling closed no zone, a roadblock and an
opening lasted until the end of the game, passive abilities were never applied,
the step granted by a sighting did not exist, and a fugitive standing still on a
respite recovered there again every time it came back.

**Stamina can be recovered, at four places on the board.** A respite is a block
in the same format as an extraction zone, set in the ring between the starting
core and the exits. The fugitive regains two points there at the end of a turn,
and the whole place goes into an eight-turn cooldown — both sides seeing which
one was used and when it returns. He therefore pays for those points in
information, as he does for every other expense.

**The street grid no longer tightens on the last block.** The fourth edge was
added as an avenue without checking its distance to the last drawn axis.

**Capture enters the game.** An inspector who stays in contact with the fugitive
across two resolutions catches him. In open ground the fugitive breaks contact
with a diagonal step; in a one-cell corridor he cannot, and the double step
becomes a survival cost again. An inspector standing on his cell now counts as
contact — refusing him that move would have told him where the fugitive is.

**A zone is no longer neutralised, its cells are occupied.** A piece posted on
one only removes an entry cell: closing a zone takes as many pieces as it has
street cells. Strangling stops with three zones left open, and its cadence
follows the length of the game instead of being fixed.

**When two end conditions fall together, the fugitive's wins.** Inspectors have
four of them, he has one.

**The starting core follows the board size, and inspectors can no longer be
placed inside it.** Its radius was five whatever the city, while sight range
followed the board's side: five pieces set around the centre spotted the
fugitive on the very first turn in nearly every game, and all the more easily
on a large board. `Settings` therefore carries a `centre_radius`, which bots
receive in `hello`.

**Appearance plugins now load and validate.** A plugin declares what it
replaces; everything else falls back on shipped content. A shape overflowing its
template, a polygon outside its three-to-thirty-two vertices, a colour missing
from the palette, a hexadecimal value written inside a shape, a prism outside
the building role: each stops the load naming the full path of the offending
key, and all are listed at once.

The check applies **at startup**, local plugins included, not only to
`filature validate` — a shape that overflows hides neighbouring tiles, which is
a gameplay advantage dressed up as decoration.

Two plugins redefining the same shape are a conflict, named on both sides.
Overriding shipped content is not.

**`filature preview <folder>` renders an appearance plugin as SVG**, without
launching the game: the sheet of its shapes on every ground, and a
board in situation. The plugin is merged onto shipped content before rendering,
and the sheet marks what comes from it — a misspelled key passes validation
while overriding nothing, and nothing else would say so.

**The `js/wasm` archive is no longer published.** It held only the `.wasm`,
unusable without `wasm_exec.js` and an HTML page. The target is still built:
it is the one check that keeps a dependency from pulling in cgo unnoticed.

**Archives now carry `THIRD-PARTY-NOTICES`**, alongside `LICENSE` and `NOTICE`.
The licences of the libraries linked into the binary require it on
redistribution, and earlier releases did not comply. The file is generated from
the actual dependencies; a test fails when it goes stale.

## [0.5.0] — 2026-08-26 — La boucle de jeu

Étape 5 de la feuille de route. **Le binaire joue** : une partie se déroule
entièrement en texte, du placement à la fin, contre un adversaire qui choisit au
hasard parmi les coups légaux. Jusqu'ici le jeu se vérifiait en test ; il se
vérifie maintenant en jouant, et une règle a déjà bougé pour cette raison.

**`shapes_version` passe à 3 : les plugins d'apparence publiés sont à
reprendre.** Deux couleurs entrent au socle obligatoire — `backdrop`, ce qu'on
voit autour du plateau, et `marker_outline`, le contour des quatre marqueurs.
Une palette qui ne les déclare pas est refusée.

Le moteur pose désormais un liseré clair sous le contour des pions et des
marqueurs, et encadre toute épaisseur de contour — jamais moins d'un pixel
d'écran, jamais plus d'un sixième de la plus petite dimension du trait. Rien à
déclarer, rien à retirer. En dessous de 24 pixels par case, le rendu ne garantit
plus la lisibilité des pions et la vue défile plutôt que de réduire.

Les faces d'un bâtiment sont dérivées de sa couleur par trois coefficients fixes
— 1,50 pour le dessus, 1,14 et 0,72 pour les côtés — que le contrat énonce.

**La validation de ce contrat n'est pas encore écrite** : `Validate` et `Merge`
sont l'étape 6. Un plugin d'apparence peut donc être accepté aujourd'hui et
refusé demain — `docs/contrat-formes.md` fait foi, pas ce que le binaire laisse
passer.

**Le jeu cherche ses extensions dans un dossier `plugins`**, à côté de
l'exécutable, et non plus dans `greffons` — le renommer suffit. Le drapeau
`--plugins` remplace `--greffons`.

Les deux sous-commandes suivent : `filature examples` et `filature validate`,
là où l'on écrivait `exemples` et `valide`.

**Tout passe à l'anglais, contrats publics compris.** Les trois numéros
l'annoncent, et changent de nom au passage : `version_effets` devient
`effects_version`, `version_formes` devient `shapes_version`, `protocole`
devient `protocol`. `effects_version` et `protocol` valent 2, `shapes_version`
vaut 3. Un plugin écrit contre une version antérieure est refusé au chargement
plutôt qu'appliqué de travers.

Ce qu'un auteur de plugin doit reprendre : les clés de son manifeste — `nom`
devient `name`, `regles` devient `rules`, `capacite` devient `ability` — et les
valeurs qu'il y écrit — `"deplacer"` devient `"step"`, `"phase_inspecteurs"`
devient `"inspectors_phase"`. Un bot doit relire `schemas/view.schema.json`,
dont toutes les clés changent. Les quatre schémas sont renommés en conséquence.

**Une graine ne rend plus le même plateau.** Les flux d'aléa portent des noms
anglais, et ces noms entrent dans le tirage : la ville de la graine 7 n'est plus
celle d'hier. Aucune partie enregistrée n'existait, mais c'est la dernière
version où ce changement était sans conséquence.

Le mot « greffon » disparaît du projet au profit de « plugin », partout — code,
documentation et libellés. Il est exclusivement français, absent de l'usage
courant, et il ne se cherche pas.

**Ce qu'un auteur de plugin d'apparence doit reprendre.** Deux couleurs entrent
au socle obligatoire : `backdrop`, ce qu'on voit autour du plateau, et
`marker_outline`, le contour des quatre marqueurs. Une palette qui ne les
déclare pas est refusée.

Le moteur pose désormais un liseré clair sous le contour des pions et des
marqueurs, et encadre toute épaisseur de contour — jamais moins d'un pixel
d'écran, jamais plus d'un sixième de la plus petite dimension du trait. Rien à
déclarer, rien à retirer. En dessous de 24 pixels par case, le rendu ne garantit
plus la lisibilité des pions et la vue défile plutôt que de réduire.

Les faces d'un bâtiment sont dérivées de sa couleur par trois coefficients fixes
— 1,50 pour le dessus, 1,14 et 0,72 pour les côtés — que le contrat énonce.

### Ajouté
- **Le binaire joue.** Lancé sans argument, il donne la main aux inspecteurs
  face à un adversaire, en texte. `--side fugitive` prend l'autre rôle,
  `--side watch` regarde deux machines s'affronter, `--preset` et `--seed`
  choisissent la partie. En spectateur, chaque tour s'affiche ; `--delay`
  espace les tours et redessine à la même place plutôt que de les empiler.
- Un adversaire qui choisit au hasard parmi les coups légaux, en
  attendant l'IA des étapes 9 et 10. Il est déterministe : une graine donnée
  rejoue la même partie.
- Chargement des plugins : capacités, dépenses et modes entrent au registre
  depuis les manifestes. Le contenu livré emprunte le même chemin qu'un plugin
  posé sur le disque, si bien qu'il est exercé à chaque démarrage.
- Les refus de `docs/plugins.md` §9 sont appliqués : un champ inconnu, une
  primitive que ce binaire ne sait pas appliquer, `rules = false` accompagné
  d'une capacité, un `differer` imbriqué, un bot mêlé à des effets. Chacun
  arrête le chargement entier en nommant le fichier et la clé fautive.
- Empreinte de contenu par plugin, qui distingue deux plugins se disant
  identiques sans l'être.
- `filature validate <dossier>`, annoncée à la version 0.2.0 : le même contrôle
  que le chargement, tous les manquements listés d'un coup avec leur fichier et
  leur chemin de clé, et l'empreinte quand le plugin tient.

### Corrigé
- La vue du fugitif annonçait une zone scellée numéro -1 tant qu'il n'avait pas
  choisi. Le champ reste maintenant absent, ce qui est ce qu'il faut lire.
- Le fugitif pouvait dépenser son dernier point de résistance, ce qui le faisait
  perdre sur-le-champ sans que l'effet acheté serve jamais. La règle lui en fait
  garder un.

***

Step 5 of the roadmap. **The binary plays**: a game runs through entirely in
text, from setup to the end, against an opponent picking at random among the
legal moves. Until now the game was checked in tests; it is now checked by
playing, and one rule has already moved because of it.

**`shapes_version` goes to 3: published appearance plugins must be revisited.**
Two colours join the required set — `backdrop`, what surrounds the board, and
`marker_outline`, the outline of all four markers. A palette declaring neither
is refused.

The engine now lays a light rim beneath the outline of pieces and markers, and
bounds every outline width — never below one screen pixel, never above one sixth
of the stroke's smallest dimension. Nothing to declare, nothing to remove. Below
24 pixels per tile, rendering no longer guarantees piece legibility and the view
scrolls rather than shrinking further.

A building's faces are derived from its colour by three fixed coefficients —
1.50 for the top, 1.14 and 0.72 for the sides — which the contract states.

**Validation of that contract is not written yet**: `Validate` and `Merge` are
step 6. An appearance plugin may therefore be accepted today and refused
tomorrow — `docs/contrat-formes.md` is what holds, not what the binary lets
through.

**The game looks for extensions in a `plugins` folder**, next to the executable,
rather than in `greffons` — renaming it is enough. The `--plugins` flag replaces
`--greffons`.

Both subcommands follow: `filature examples` and `filature validate`, where one
used to write `exemples` and `valide`.

**Everything moves to English, public contracts included.** The three numbers
say so, and are renamed on the way: `version_effets` becomes `effects_version`,
`version_formes` becomes `shapes_version`, `protocole` becomes `protocol`.
`effects_version` and `protocol` are now 2, `shapes_version` is 3. A plugin
written against an earlier version is refused at load rather than applied
wrongly.

What a plugin author must revisit: their manifest keys — `nom` becomes `name`,
`regles` becomes `rules`, `capacite` becomes `ability` — and the values written
there — `"deplacer"` becomes `"step"`, `"phase_inspecteurs"` becomes
`"inspectors_phase"`. A bot must re-read `schemas/view.schema.json`, whose every
key changes. All four schemas are renamed accordingly.

**A seed no longer yields the same board.** Random streams carry English names,
and those names feed the draw: seed 7's city is not yesterday's. No saved game
existed, but this is the last release where that change came free.

The word « greffon » is gone from the project in favour of « plugin », everywhere —
code, documentation and labels. It is exclusively French, absent from common
usage, and nobody searches for it.

**What an appearance plugin author must revisit.** Two colours join the required
set: `backdrop`, what surrounds the board, and `marker_outline`, the outline of
all four markers. A palette that declares neither is refused.

The engine now lays a light rim beneath the outline of pieces and markers, and
bounds every outline width — never below one screen pixel, never above one sixth
of the stroke's smallest dimension. Nothing to declare, nothing to remove. Below
24 pixels per tile, rendering no longer guarantees piece legibility and the view
scrolls rather than shrinking further.

A building's faces are derived from its colour by three fixed coefficients —
1.50 for the top, 1.14 and 0.72 for the sides — which the contract states.

### Added
- **The binary plays.** Run with no argument, it hands you the inspectors
  against an opponent, in text. `--side fugitive` takes the other side,
  `--side watch` watches two machines play, `--preset` and `--seed` choose
  the game. When watching, every turn is shown; `--delay` spaces them out and
  redraws in place rather than stacking them.
- An opponent drawing its moves at random among the legal ones, standing in
  until the AI of steps 9 and 10. It is deterministic: a given seed replays the
  same game.
- Plugin loading: abilities, expenses and modes now enter the registry from
  manifests. Shipped content takes the same path as a plugin dropped on disk,
  so that path is exercised on every start.
- The refusals in `docs/plugins.md` §9 are enforced: an unknown field, a
  primitive this binary cannot apply, `rules = false` alongside an ability, a
  nested `differer`, a bot mixed with effects. Each aborts the whole load,
  naming the file and the offending key.
- Per-plugin content fingerprint, telling apart two plugins that claim to be
  identical without being so.
- `filature validate <folder>`, announced back in 0.2.0: the same checks the
  loader runs, every failure listed at once with its file and key path, and the
  fingerprint when the plugin holds.

### Fixed
- The fugitive's view reported a sealed zone numbered -1 until he had chosen
  one. The field is now simply absent, which is what should be read.
- The fugitive could spend his last stamina point, losing on the spot without
  the purchased effect ever serving. The rule makes him keep one.

## [0.4.0] — 2026-08-26 — La vision

Étape 4 de la feuille de route. Le fugitif peut désormais être repéré : jusqu'à
cette version il ne l'était jamais, quoi qu'il fasse, et ni la portée de vue ni
les capacités des inspecteurs ne changeaient quoi que ce soit.

**Le binaire ne joue toujours pas** : la boucle de jeu est l'étape 5.

Rien à reprendre pour un auteur de greffon : `version_formes`, `protocole` et
`version_effets` sont inchangés. Un bot recevra en revanche des vues plus
fournies — le champ des cases couvertes n'était jamais rempli.

### Ajouté
- Vision : un inspecteur voit dans les huit directions jusqu'à sa portée, la
  ligne s'arrêtant au premier bâtiment, au premier barrage ou au premier autre
  inspecteur. Le fugitif est repéré en entrant dans une ligne de vue et cesse de
  l'être en sortant — jusqu'ici il ne l'était jamais, quoi qu'il fasse.
- La vue filtrée des inspecteurs porte les cases qu'ils couvrent réellement,
  occlusion comprise : rien de ce qui se trouve derrière un collègue ou un
  barrage n'y figure, et chacun voit la case où il se tient.

***

Step 4 of the roadmap. The fugitive can now be spotted: until this release he
never was, whatever he did, and neither sight range nor the inspectors'
abilities changed anything at all.

**The binary still does not play**: the game loop is step 5.

Nothing for plugin authors to update: `version_formes`, `protocole` and
`version_effets` are unchanged. A bot will however receive fuller views — the
covered-cells field was never populated.

### Added
- Sight: an inspector sees in all eight directions up to their range, the line
  stopping at the first building, roadblock or fellow inspector. The fugitive is
  spotted on entering a line of sight and stops being so on leaving it — until
  now he never was, whatever he did.
- The inspectors' filtered view carries the cells they actually cover, occlusion
  included: nothing behind a fellow inspector or a roadblock appears in it, and
  each of them sees the cell he stands on.

## [0.3.0] — 2026-08-25 — Le plateau

Étape 3 de la feuille de route. Le plateau se fabrique au lieu d'être fourni :
une graine donne une ville, et la même graine la redonne à l'identique sur
Windows, Linux et macOS.

**Le binaire ne joue toujours pas** : la boucle de jeu est l'étape 5.

Rien à reprendre pour un auteur de greffon : `version_formes`, `protocole` et
`version_effets` sont inchangés.

### Ajouté
- Génération de plateau : trame irrégulière, cours percées, impasses et six
  zones d'extraction réparties sur le pourtour. Une graine donnée produit
  toujours le même plateau, vérifié sur Windows, Linux et macOS.
- Trois préréglages de partie — Quartier, Faubourg et Ville — là où seule la
  Ville existait. La portée de vue et la durée suivent la taille du plateau.

### Corrigé
- Les bornes des paramètres avaient des valeurs sans nom ni justification, et
  les refus ne disaient ni la valeur reçue ni ce qui était attendu. Le nombre de
  pions déplaçables n'était pas contrôlé.
- Le README annonçait un moteur non écrit alors qu'une partie complète se joue
  depuis des appels Go.
- Le creusement des impasses n'en produisait aucune : les couloirs rejoignaient
  l'avenue suivante au lieu de rester borgnes, et le piégeage sur lequel repose
  le Barreur n'existait pas. Leur nombre suit désormais la surface du plateau et
  non son côté, qui rendait une Ville deux fois moins piégeuse qu'un Quartier.

***

Step 3 of the roadmap. The board is now built rather than supplied: a seed gives
a city, and the same seed gives it back identically on Windows, Linux and macOS.

**The binary still does not play**: the game loop is step 5.

Nothing for plugin authors to update: `version_formes`, `protocole` and
`version_effets` are unchanged.

### Added
- Board generation: irregular street grid, punched courtyards, dead ends and six
  extraction zones spread around the perimeter. A given seed always produces the
  same board, verified on Windows, Linux and macOS.
- Three game presets — District, Outskirts and City — where only the City
  existed. Sight range and duration follow the board size.

### Fixed
- Parameter bounds were unnamed magic numbers, and rejections stated neither the
  value received nor what was expected. The number of movable pieces was not
  checked at all.
- The README claimed the engine was not written yet, when a full game already
  plays out from Go calls.
- Dead-end carving produced none: corridors reached the next avenue instead of
  staying blind, and the trapping the Blocker relies on did not exist. Their
  count now follows the board's area rather than its side, which made a City
  half as trap-ridden as a District.

## [0.2.0] — 2026-08-25 — La vue filtrée

Étape 2 de la feuille de route. Les quatre invariants de l'architecture sont
désormais posés et gardés par des tests.

**Le binaire ne joue toujours pas** : la boucle de jeu est l'étape 5.

`schemas/vue.schema.json` paraît pour la première fois. Un bot peut s'écrire
contre lui — c'est exactement ce que le jeu enverra, puisque le schéma est
généré depuis la structure et qu'un test échoue s'ils divergent.

### Ajouté
- Vue filtrée : chaque camp ne reçoit que ce qu'il a le droit de savoir. La
  position du fugitif, sa zone scellée, sa résistance et les traces hors de
  portée ne franchissent pas la projection.
- `schemas/vue.schema.json`, généré depuis `noyau.Vue` et non écrit à la main.
- La validation d'un greffon avant installation est prévue pour l'étape 8 :
  `filature valide <chemin>`, avec le fichier, la ligne et le chemin de clé de
  chaque manquement.

### Corrigé
- Les listes de la vue rendaient `null` au lieu d'un tableau vide, ce qui
  obligeait un bot à traiter deux formes pour la même absence.
- Les traces se découvrent en distance de Manhattan, comme la règle l'énonce, et
  non de Tchebychev qui aurait ouvert les quatre diagonales.

***

Step 2 of the roadmap. All four architectural invariants are now in place and
guarded by tests.

**The binary still does not play**: the game loop is step 5.

`schemas/vue.schema.json` appears for the first time. A bot can be written
against it — it is exactly what the game will send, since the schema is
generated from the structure and a test fails if the two drift apart.

### Added
- Filtered view: each side receives only what it is entitled to know. The
  fugitive's position, sealed zone, stamina and out-of-range trails do not cross
  the projection.
- `schemas/vue.schema.json`, generated from `noyau.Vue` rather than written by
  hand.
- Validating a plugin before install is planned for step 8: `filature valide
  <path>`, reporting the file, line and key path of each failure.

### Fixed
- The view's lists returned `null` instead of an empty array, forcing a bot to
  handle two shapes for the same absence.
- Trails are found at Manhattan distance, as the rules state, not Chebyshev
  which would have opened the four diagonals.

## [0.1.1] — 2026-08-25 — Correctifs du binaire

### Corrigé
- Le dossier des greffons est cherché à côté de l'exécutable et non dans le
  répertoire courant. Lancé par un raccourci, le jeu ignorait les greffons
  installés sans rien dire.
- `filature exemples` refuse d'écrire dans le dossier des greffons actifs. Le
  contenu livré y aurait été déclaré deux fois — une fois depuis le binaire, une
  fois depuis le disque — et deux greffons qui définissent la même clé sont un
  conflit.

### Ajouté
- Le README décrit les deux commandes du binaire et l'emplacement des greffons.

***

Two fixes to the binary. The core is unchanged, and the game still does not
launch — the loop is step 5.

### Fixed
- The plugin folder is looked up next to the executable, not in the current
  directory. Launched from a shortcut, the game silently ignored installed
  plugins.
- `filature exemples` refuses to write into the active plugin folder. Shipped
  content would have been declared twice — once from the binary, once from
  disk — and two plugins defining the same key are a conflict.

### Added
- The README documents the binary's two commands and where plugins live.

## [0.1.0] — 2026-08-25 — Le noyau

Étape 1 de la feuille de route : le noyau. Une partie complète se joue depuis
des appels Go, sans interface.

**Le binaire ne joue pas encore.** `filature` lancé répond qu'il reste à
implémenter : la boucle de jeu est l'étape 5. Cette version marque un jalon de
développement, pas une version jouable.

Rien à reprendre pour un auteur de greffon : `version_formes`, `protocole` et
`version_effets` valent 1, et c'est leur première publication.

### Ajouté
- Squelette du projet, contrats de formes et de bot.
- Vocabulaire d'effets documenté, avec les primitives `differer` et
  `ouvrir_case`.
- File des effets différés dans l'état, et les différés annoncés dans la vue.
- Modes de jeu déclaratifs, et `version_effets` au manifeste.
- Étranglement exprimé en effets dans `greffons/base`, plus codé en dur.
- Documentation du format d'un greffon, et table `bot` au schéma de manifeste.
- Application et annulation des dix-sept primitives d'effets.
- Greffons de langue, avec le français en repli et l'anglais livré.
- Contenu livré embarqué dans le binaire, et `filature exemples` pour l'en
  ressortir.
- Énumération des coups légaux, et la dépense de changement de zone qui
  manquait au contenu livré.
- Application et annulation d'un coup, avec l'enchaînement des phases.
- Générateur déterministe à flux nommés, seul aléatoire du jeu.
- Résolution de fin de tour : contacts plafonnés, traces, révélation
  périodique, effets différés et étranglement.
- Arbitrage : extraction sur deux tours, zone neutralisée par un inspecteur ou
  par sa fermeture, et les trois victoires des inspecteurs.
- Mise en place d'une partie, plateau reçu et non fabriqué.

***

Step 1 of the roadmap: the core. A full game plays out from Go calls, with no
interface.

**The binary does not play yet.** Running `filature` reports that it remains to
be implemented: the game loop is step 5. This release marks a development
milestone, not a playable version.

Nothing for plugin authors to update: `version_formes`, `protocole` and
`version_effets` are all 1, published here for the first time.

### Added
- Project skeleton, shape and bot contracts.
- Effect vocabulary documented, with the `differer` and `ouvrir_case`
  primitives.
- Deferred effect queue in the state, and announced deferrals in the view.
- Declarative game modes, and `version_effets` in the manifest.
- Strangling expressed as effects in `greffons/base`, no longer hardcoded.
- Documentation of a plugin's format, and a `bot` table in the manifest schema.
- Apply and undo for the seventeen effect primitives.
- Language plugins, with French as fallback and English shipped.
- Shipped content embedded in the binary, and `filature exemples` to write it
  back out.
- Legal move enumeration, and the zone-change expense missing from shipped
  content.
- Apply and undo for a move, with phase sequencing.
- Deterministic generator with named streams, the game's only randomness.
- End-of-turn resolution: capped contacts, trails, periodic reveal, deferred
  effects and strangling.
- Arbitration: extraction over two turns, zone neutralised by an inspector or by
  its closing, and the inspectors' three victories.
- Game setup, with the board received rather than built.
