# Cibles en anglais, commentaires en français : convention des autres projets.

BINAIRE   ?= evasion
VERSION   ?= dev
LDFLAGS    = -s -w -X main.version=$(VERSION)
SORTIE    ?= .tmp

# makefile.local porte ce qui est propre au poste — GOTMPDIR, GOCACHE — et n'est
# pas versionné. Inclus ici et non en tête : une affectation immédiate qui y
# référencerait une variable définie plus haut trouverait une chaîne vide.
#
# Passer par le Makefile plutôt que d'appeler go à la main : une commande tapée
# directement perd ces variables, et l'échec est intermittent.
-include makefile.local

.PHONY: build run test race lint vulncheck sec cover binaries clean tools

build:
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(SORTIE)/$(BINAIRE) ./cmd/evasion

run:
	go run ./cmd/evasion

# Aucun test ne doit ouvrir de fenêtre : les runners sont sans écran, et un test
# qui exigerait xvfb n'a rien à faire dans la suite par défaut.
#
# PKG et RUN restreignent la portée sans sortir du Makefile : une commande go
# tapée directement perd les réglages de makefile.local, et l'oubli ne se voit
# pas dans la sortie.
PKG ?= ./...
RUN ?=

# Le compte des mécaniques inertes est affiché avec la suite, jamais séparément :
# un test qui se saute passe inaperçu là où un test rouge fait mal, et six mois
# plus tard il reste des skips que personne ne regarde. C'est le pendant du
# compte de stubs de ROADMAP.md — il descend tout seul et ne ment pas.
#
# Le compte vient du test, jamais d'un motif recopié ici. La première version
# posait un grep dont le motif divergeait d'un espace de celui du fichier : elle
# affichait zéro, ce qui est pire qu'aucun compteur. Aucun filtre non plus sur la
# sortie — il dépendrait du texte du message et casserait à la première
# reformulation, en silence.
test:
	go test $(if $(RUN),-run '$(RUN)') $(PKG)
	@go test -v -run TestInertMechanicsAreCounted ./internal/quality/

race:
	CGO_ENABLED=1 go test -race ./...

lint:
	golangci-lint run

vulncheck:
	govulncheck ./...

# .tmp porte les programmes jetables, qui ne sont pas versionnés : gosec les
# analyse en local et pas en intégration continue, où le dossier n'existe pas.
# Sans exclusion, la cible échoue ici et passe là-bas, ce qui finit par la faire
# ignorer — et par masquer un vrai signalement.
sec:
	gosec -exclude-dir=.tmp ./...

# Les licences des dépendances liées au binaire. Un test compare le fichier aux
# dépendances réelles et échoue quand il vieillit : la cible ne sert qu'à le
# réécrire, et son diff se relit comme le reste.
notices:
	go test -run TestThirdPartyNoticesAreCurrent ./internal/quality/ -maj-notices

cover:
	go test -coverprofile=$(SORTIE)/couverture.out ./...
	go tool cover -html=$(SORTIE)/couverture.out

# Une seule cible paramétrée, pas cinq recettes : la matrice complète vit dans
# le workflow, qui est le seul endroit où elle peut être exécutée. En local on
# ne produit que ce qui se croise sans compilateur C — Windows et wasm.
binary:
	CGO_ENABLED=$(CGO) GOOS=$(OS) GOARCH=$(ARCH) go build -trimpath \
	  -ldflags "$(LDFLAGS)" -o dist/$(BINAIRE)_$(OS)_$(ARCH)$(EXT) ./cmd/evasion

binaries:
	$(MAKE) binary OS=windows ARCH=amd64 CGO=0 EXT=.exe
	$(MAKE) binary OS=js ARCH=wasm CGO=0
	@echo "Les cibles linux et darwin exigent une compilation native : voir release.yml"

clean:
	rm -rf $(SORTIE) dist

# Les versions de golangci-lint et de gosec sont épinglées et partagées avec le
# workflow, qui les reprend telles quelles. Avec @latest des deux côtés, le
# poste prend du retard dès qu'une version sort : le contrôle passe en local et
# échoue en intégration continue, sur du code que personne n'a touché. Un
# numéro changé ici se change donc aussi dans le workflow.
#
# L'un comme l'autre lit les données d'export du compilateur par le
# golang.org/x/tools qu'il embarque, et doit donc en porter une version qui
# connaisse le format produit par le Go des runners. En deçà, chaque import de
# chaque paquet échoue sur « export data version N is greater than maximum
# supported », et les deux contrôles tombent sans qu'un fichier du dépôt ait
# bougé.
#
# La ligne de Go ne se borne pas pour autant, et c'est tranché. L'épingler
# priverait le binaire des correctifs du runtime, et fausserait govulncheck, qui
# évalue ses avis contre la bibliothèque standard du toolchain : c'est
# l'outillage qui suit Go, jamais l'inverse. Retirer check-latest du workflow
# n'y changerait rien non plus — quand l'image des runners ne porte pas la ligne
# que go.mod demande, setup-go va chercher la plus haute de cette ligne de toute
# façon.
GOLANGCI_VERSION ?= v2.14.0

# gosec n'a pas de release portant un x/tools qui lise les données d'export de
# Go 1.27.2 : v2.29.0 embarque la v0.49.0, et le commit qui monte à la v0.51.0
# n'est pas tagué. D'où ce commit plutôt qu'une version — à remplacer par la
# release dès qu'elle sort.
GOSEC_VERSION ?= v2.29.1-0.20261009120814-7b1b5cebe007

tools:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_VERSION)
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/securego/gosec/v2/cmd/gosec@$(GOSEC_VERSION)
