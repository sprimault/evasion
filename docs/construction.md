# Construction

Le binaire est le mode d'installation : quelqu'un télécharge un fichier et
l'exécute, sans Go, sans runtime, sans dépendance système. Tout ce qui suit en
découle.

---

## La matrice

**Une seule machine ne produit pas toutes les cibles.** Ebitengine ne se compile
en croisé que vers Windows et WebAssembly ; Linux exige une compilation native,
parce que le moteur y passe par les API graphiques du système.

| Cible | Où | cgo | Publiée |
|---|---|---|---|
| windows/amd64 | Linux, croisé | `0` | oui |
| js/wasm | Linux, croisé | `0` | **non** |
| linux/amd64 | Linux, natif | `1` | oui |
| linux/arm64 | Linux ARM, natif | `1` | oui |

**macOS n'est plus une cible.** Aucun runner hébergé n'acquiert ses jobs : ils
attendent un quart d'heure, puis sont annulés sur « not acquired by Runner of
type hosted even after multiple attempts ». La cible n'est donc ni compilée ni
testée, et un binaire que rien ne vérifie ne se publie pas — pire, un contrôle
qui ne rend jamais bloque la fusion de tout le reste.

Le noyau reste portable, et rien dans le code ne suppose un système : c'est la
**publication** qui s'arrête, pas le support. Une cible macOS redevient une ligne
de matrice le jour où un runner la prend.

**`js/wasm` se compile sans être publiée**, et ce n'est pas un oubli. L'archive
serait inutilisable seule — il y faudrait `wasm_exec.js` et une page HTML — et le
mode web n'est pas au programme : un plugin tiers y arriverait par le navigateur
plutôt que par un dossier, et un bot externe, qui est un processus, n'y arrive
pas du tout.

La cible reste dans la matrice parce qu'elle **empêche une dépendance
d'introduire du cgo** sans que personne le voie — `windows/amd64`, compilé sans
cgo lui aussi, attrape la même chose. Elle a d'autres raisons d'y être : la
surface d'API de `js` et l'absence d'appels système y sont vérifiées, et nulle
part ailleurs.

## cgo n'est pas uniforme

C'est la particularité du projet, et elle se contourne sans bruit : rien
n'avertit qu'on vient de la perdre.

Une valeur globale est donc proscrite : `cgo` est une variable de matrice, et un
`0` posé là où le moteur attend les API du système ne produit pas une erreur
claire mais un échec de linkage obscur.

Conséquence à connaître : le binaire Linux est lié à la glibc. Il n'y a pas de
construction statique musl, donc pas d'image `scratch` possible — sans
importance pour un jeu de bureau.

**Aucune autre dépendance n'a le droit d'introduire cgo.** `modernc.org/sqlite`
et `wazero` sont choisis pour ça : une seconde source rendrait les cibles
Windows et WebAssembly impossibles.

## Le Makefile ne porte pas la matrice

`make binary` prend `OS`, `ARCH` et `CGO` en paramètres. La liste des couples
vit dans les workflows, seuls endroits où elle s'exécute réellement.

**Elle y est écrite deux fois**, dans l'intégration et dans la publication, et
la duplication est assumée : une matrice partagée devrait passer par un workflow
appelé, ce qui coûterait plus que ce qu'elle éviterait. Les deux ont déjà
divergé une fois sur le nombre de cibles publiées.

En local, `make binaries` ne produit que ce qui se croise — Windows et
WebAssembly — et le dit.

## Options de compilation

```
CGO_ENABLED=$CGO GOOS=$OS GOARCH=$ARCH go build -trimpath \
  -ldflags "-s -w -X main.version=$VERSION" -o dist/evasion_${OS}_${ARCH}
```

`-trimpath` toujours : les chemins absolus de la machine de construction n'ont
rien à faire dans un artefact public.

Sous Windows, `-H windowsgui` supprime la console au lancement — mais alors
`evasion version` n'écrit nulle part, la sortie standard n'existant pas. Le
contrôle de version de la chaîne de publication en dépend : soit on renonce au
drapeau, soit on attache une console à la demande.

## Empaquetage

`.zip` pour Windows, `.tar.gz` ailleurs : c'est ce que chaque système ouvre sans
rien installer, et le tar préserve le bit exécutable qu'un binaire téléchargé nu
perdrait.

`LICENSE` et `NOTICE` vont **dans** l'archive et non à côté : Apache 2.0 les
exige à la redistribution, et personne ne télécharge un fichier séparé pour
accompagner un binaire.

`THIRD-PARTY-NOTICES` les accompagne, au titre des licences des bibliothèques
liées au binaire. La MIT, comme la plupart des licences permissives, exige que
sa notice suive « toute copie ou partie substantielle » du logiciel — un binaire
qui la compile en est une.

**Ce fichier se génère, il ne se tient pas à la main** : `make notices` le
réécrit depuis les dépendances que `go list -deps` trouve dans la commande, donc
sans l'outillage ni les dépendances de test. Un test le compare aux dépendances
réelles à chaque exécution de la suite et échoue quand il vieillit, faute de quoi
une bibliothèque ajoutée entrerait dans les archives sans sa notice, et personne
ne s'en apercevrait.

**Il porte l'union des cibles de la matrice, pas ce que lie la machine qui le
produit.** Le jeu de modules dépend du système — un pilote de fenêtre n'existe
que sur Linux, un paquet système disparaît sur WebAssembly. Un fichier par cible
serait plus juste au sens strict ; l'union est retenue parce qu'un lecteur reçoit
alors la même liste quelle que soit son archive, et qu'une notice de trop n'a
jamais fait de tort à personne. Un test compare la liste des cibles à celle du
workflow de publication, sans quoi les deux divergeraient.

**L'archive ne contient rien d'autre que le binaire et ces trois fichiers.** Le
contenu livré — règles, formes, palette, français et anglais — est embarqué dans
l'exécutable par `//go:embed` : un binaire déplacé continue de fonctionner, et
personne ne casse ses règles en éditant un fichier qu'il n'a pas écrit.

Ce contenu reste accessible à qui veut s'en inspirer : `evasion examples
<dossier>` l'écrit sur le disque, sans jamais écraser un fichier existant. C'est
par là que passe un traducteur qui recopie l'anglais pour en faire sa langue.

Le nom de l'archive porte version, système et architecture ; le binaire qu'elle
contient s'appelle `evasion` tout court — une fois extrait, on sait sur quelle
machine on est.

Trois opérations ne peuvent pas se faire cible par cible et vont dans une étape
de rassemblement : **`SHA256SUMS`, calculé en un seul endroit sur toutes les
archives**, l'attestation de provenance, et la création de la version.

Les sommes portent sur les archives et non sur les binaires : c'est l'archive
que la personne télécharge, donc ce qu'elle peut vérifier.

## Signature

Le binaire Windows n'est pas signé : SmartScreen affichera un avertissement au
premier lancement. C'est dit dans le README plutôt que laissé à découvrir.

L'attestation de provenance lie chaque archive au commit et à la chaîne qui l'a
produite, et se vérifie avec `gh attestation verify`. C'est le seul contrepoids
crédible à des binaires non signés.

## Pas d'image Docker

Le mode réseau est de pair à pair entre deux joueurs : personne ne déploie un
conteneur pour faire une partie, et le binaire ouvre une fenêtre.

Un serveur dédié sans écran serait un binaire séparé, avec une étiquette de
compilation excluant le rendu. Le noyau n'importe rien de graphique, donc c'est
faisable — mais rien ne le réclame aujourd'hui.

## Tests et intégration continue

Les machines d'intégration sont sans écran. `internal/core` et `internal/ai`
n'importent pas Ebitengine, ce qui rend la contrainte facile à tenir : aucun
test de la suite par défaut n'ouvre de fenêtre.

Un test qui exigerait `xvfb-run` sous Linux est le signe qu'il n'a pas sa place
dans cette suite. Le rendu se vérifie à l'œil.

`make race` demande cgo, y compris sous Windows où Ebitengine s'en passe : sans
compilateur C installé, le message d'erreur parle de cgo et non du détecteur de
courses, ce qui envoie chercher au mauvais endroit.
