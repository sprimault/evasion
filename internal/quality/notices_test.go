// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

package quality

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"
	"testing"
)

// majNotices réécrit les notices de tiers au lieu de les comparer.
//
// Jamais automatique, pour la même raison que les autres attendus : un fichier
// régénéré sans être relu ne prouve plus rien, et celui-ci porte des obligations
// juridiques.
var majNotices = flag.Bool("maj-notices", false, "réécrire THIRD-PARTY-NOTICES")

// cheminNotices est le fichier livré dans chaque archive, à côté de LICENSE.
var cheminNotices = filepath.Join(racine, "THIRD-PARTY-NOTICES")

// enTeteNotices explique le fichier à qui l'ouvre, et pourquoi il existe.
const enTeteNotices = `Notices des logiciels tiers
===========================

Le binaire Evasion incorpore les bibliothèques ci-dessous. Leurs licences
exigent que ces notices accompagnent toute redistribution, forme binaire
comprise.

Ce fichier est généré depuis les dépendances réellement liées au binaire :
« make notices ». Ne pas le modifier à la main.

Chaque entrée porte son identifiant SPDX, reconnu au texte de la licence et non
déclaré à la main. Un module offert sous plusieurs licences au choix n'engage
qu'à l'une d'elles, et la ligne de déclaration dit donc laquelle est exercée :
c'est Apache-2.0, celle du dépôt.

`

// TestThirdPartyNoticesAreCurrent vérifie que les notices livrées couvrent
// exactement ce que le binaire incorpore.
//
// La licence MIT de BurntSushi/toml, comme la plupart des licences permissives,
// exige que sa notice accompagne « toute copie ou partie substantielle » du
// logiciel — un binaire qui la compile en est une. Sans ce test, une dépendance
// ajoutée entrerait dans les archives sans sa notice, et personne ne s'en
// apercevrait avant que quelqu'un le remarque.
func TestThirdPartyNoticesAreCurrent(t *testing.T) {
	attendu, err := notices()
	if err != nil {
		t.Fatal(err)
	}

	if *majNotices {
		if err := os.WriteFile(cheminNotices, []byte(attendu), 0o600); err != nil {
			t.Fatal(err)
		}
		t.Log("THIRD-PARTY-NOTICES réécrit — relire le diff")
		return
	}

	actuel, err := os.ReadFile(cheminNotices)
	if err != nil {
		t.Fatalf("%v — lancer « make notices »", err)
	}
	if string(actuel) != attendu {
		t.Error("les notices de tiers ne correspondent plus aux dépendances du binaire — " +
			"lancer « make notices » et relire le diff")
	}
}

// notices assemble le fichier depuis les modules réellement liés au binaire.
func notices() (string, error) {
	modules, err := modulesLies()
	if err != nil {
		return "", err
	}

	var b strings.Builder
	b.WriteString(enTeteNotices)

	for _, m := range modules {
		texte, jeton, err := licenceDe(m.dossier)
		if err != nil {
			return "", fmt.Errorf("%s: %w", m.chemin, err)
		}
		fmt.Fprintf(&b, "-------------------------------------------------------------------------------\n%s %s — %s\n%s\n\n%s\n",
			m.chemin, m.version, jeton, m.origine, strings.TrimRight(texte, "\n"))
	}
	return b.String(), nil
}

// module décrit une dépendance liée au binaire.
type module struct {
	chemin  string
	version string
	dossier string
	origine string
}

// cible est un couple de la matrice de publication.
type cible struct{ os, arch, cgo string }

// ciblesPubliees suit la matrice de .github/workflows/release.yml, js/wasm
// compris : il ne se publie pas, mais il est ce qui empêche une dépendance
// d'introduire du cgo sans qu'on le voie, et ses modules liés sont les plus
// pauvres — les omettre ne changerait rien, les garder coûte une seconde.
//
// Un test compare cette liste au workflow : deux définitions finiraient par
// diverger, et c'est celle du workflow qu'on ne peut pas essayer avant de
// pousser.
var ciblesPubliees = []cible{
	{"windows", "amd64", "0"},
	{"js", "wasm", "0"},
	{"linux", "amd64", "1"},
	{"linux", "arm64", "1"},
}

// modulesLies interroge le compilateur plutôt que go.mod, pour chaque cible
// publiée.
//
// go.mod porte aussi ce que seuls les tests et l'outillage utilisent ; ce qui
// compte ici est ce que le binaire publié incorpore, et « go list -deps » sur la
// commande le donne exactement.
//
// **L'union des cibles, et non celle de la machine qui exécute.** Le jeu de
// modules liés dépend du système : un pilote de fenêtre n'existe que sur Linux,
// un paquet système disparaît sur WebAssembly. Sans GOOS fixé, le fichier valait
// pour le poste qui l'a produit, la suite rougissait sur les trois autres
// runners, et l'archive d'une plateforme serait partie sans la notice de ce
// qu'elle seule incorpore — soit l'obligation même que ce fichier existe pour
// tenir.
//
// Rien ne peut encore l'éprouver : la seule dépendance du binaire est en Go pur
// et rend la même liste sur les quatre cibles. C'est la dépendance suivante qui
// fera la différence, et c'est pour elle que le mécanisme est posé maintenant.
func modulesLies() ([]module, error) {
	vus := map[string]module{}
	for _, c := range ciblesPubliees {
		if err := listerPour(c, vus); err != nil {
			return nil, err
		}
	}

	if len(vus) == 0 {
		return nil, fmt.Errorf("aucune dépendance trouvée, ce qui n'arrive pas : la commande a-t-elle compilé ?")
	}

	chemins := make([]string, 0, len(vus))
	for c := range vus {
		chemins = append(chemins, c)
	}
	sort.Strings(chemins)

	modules := make([]module, 0, len(chemins))
	for _, c := range chemins {
		modules = append(modules, vus[c])
	}
	return modules, nil
}

// listerPour ajoute à vus les modules que le binaire lie sur une cible.
//
// Une même dépendance vue sur deux cibles y entre une fois : c'est le même
// module à la même version, et son fichier de licence est le même.
func listerPour(c cible, vus map[string]module) error {
	commande := exec.Command("go", "list", "-deps",
		"-f", "{{if .Module}}{{.Module.Path}}\t{{.Module.Version}}\t{{.Module.Dir}}{{end}}",
		"github.com/sprimault/evasion/cmd/evasion")
	commande.Env = append(os.Environ(),
		"GOOS="+c.os, "GOARCH="+c.arch, "CGO_ENABLED="+c.cgo)

	sortie, err := commande.Output()
	if err != nil {
		return fmt.Errorf("liste des dépendances pour %s/%s : %w", c.os, c.arch, err)
	}

	for _, ligne := range strings.Split(strings.ReplaceAll(string(sortie), "\r\n", "\n"), "\n") {
		champs := strings.Split(strings.TrimSpace(ligne), "\t")
		if len(champs) != 3 || champs[0] == "" {
			continue
		}
		// Le module du jeu lui-même porte sa propre licence, dans LICENSE.
		if champs[0] == "github.com/sprimault/evasion" {
			continue
		}
		vus[champs[0]] = module{
			chemin:  champs[0],
			version: champs[1],
			dossier: champs[2],
			origine: "https://" + champs[0],
		}
	}
	return nil
}

// TestNoticeTargetsMatchTheWorkflow vérifie que les cibles interrogées sont
// celles que la publication construit.
//
// Deux listes qui décrivent la même matrice divergent, et c'est celle du
// workflow qu'on ne peut pas essayer avant de pousser : une cible ajoutée là-bas
// et oubliée ici ferait partir son archive sans les notices de ce qu'elle seule
// incorpore.
func TestNoticeTargetsMatchTheWorkflow(t *testing.T) {
	brut, err := os.ReadFile(filepath.Join(racine, ".github", "workflows", "release.yml"))
	if err != nil {
		t.Fatal(err)
	}

	ligne := regexp.MustCompile(`\{\s*os:\s*(\w+),\s*arch:\s*(\w+),\s*cgo:\s*(\d)`)
	var publiees []cible
	for _, trouvaille := range ligne.FindAllStringSubmatch(string(brut), -1) {
		publiees = append(publiees, cible{trouvaille[1], trouvaille[2], trouvaille[3]})
	}
	if len(publiees) == 0 {
		t.Fatal("aucune cible lue dans release.yml : le contrôle ne dirait rien")
	}

	if !slices.Equal(publiees, ciblesPubliees) {
		t.Errorf("release.yml construit %v, les notices interrogent %v", publiees, ciblesPubliees)
	}
}

// licenceExercee est la licence retenue quand un module en offre plusieurs au
// choix.
//
// Celle du dépôt, pour que tout ce qu'une archive porte relève d'un seul régime.
// Le choix appartient à qui redistribue, et les licences d'une double offre ne
// se cumulent pas : en exercer une suffit, en recopier deux sans dire laquelle
// laisse le lecteur décider à notre place.
const licenceExercee = "Apache-2.0"

// nomsLicence est la liste close des noms de fichier de licence reconnus.
//
// Close et non un motif « LICENSE* » : un module peut poser à côté de sa licence
// un fichier qui n'en est pas une — les notices de ses propres dépendances, par
// exemple —, et un motif le ramasserait comme s'il engageait le module.
var nomsLicence = []string{
	"LICENSE", "LICENSE.txt", "LICENSE.md",
	"COPYING", "COPYING.txt",
	"LICENCE", "LICENCE.txt",
	"LICENSE-APACHE", "LICENSE-APACHE.txt",
	"LICENSE-MIT", "LICENSE-MIT.txt",
}

// licenceConnue associe un identifiant SPDX aux fragments qui le reconnaissent.
type licenceConnue struct {
	jeton     string
	marqueurs []string
}

// licencesConnues reconnaît une licence à son texte.
//
// Une tranche et non une table : le parcours d'une map n'a pas d'ordre stable, et
// ce qui décide d'un jeton ne passe jamais par là.
//
// Deux fragments par licence, parce qu'un seul ne suffit pas — « Apache License »
// figure dans l'appendice de tout fichier qui explique comment l'appliquer, et
// mentionner une licence n'est pas en être une. Les deux retenus sont ceux que le
// corps normatif porte et qu'un en-tête ne porte pas : « The MIT License (MIT) »
// et « MIT License » coexistent chez deux de nos dépendances, donc l'en-tête ne
// reconnaît rien.
//
// La table ne couvre que ce que le binaire lie. C'est voulu : un texte inconnu
// arrête le contrôle au lieu d'être deviné, et un jeton faux serait pire que pas
// de jeton puisqu'il affirmerait un régime que personne n'a vérifié.
var licencesConnues = []licenceConnue{
	{"Apache-2.0", []string{"Apache License", "Version 2.0, January 2004"}},
	{"MIT", []string{"Permission is hereby granted, free of charge", `THE SOFTWARE IS PROVIDED "AS IS"`}},
}

// licenceDe rend le texte de licence d'un module et son identifiant SPDX.
//
// Ne pas en trouver est une erreur et non un cas à ignorer : une dépendance sans
// licence lisible ne peut pas être redistribuée en confiance.
//
// **Un module qui en offre plusieurs au choix est le cas qui demande une
// décision**, et c'est [licenceExercee] qui la porte. Si ce qu'il offre ne la
// comprend pas, l'erreur nomme ce qu'il propose : choisir à la place de
// l'intégrateur engagerait la redistribution sur un régime que rien n'aurait
// arbitré.
func licenceDe(dossier string) (texte, jeton string, err error) {
	offertes := map[string]string{}
	for _, nom := range nomsLicence {
		brut, err := os.ReadFile(filepath.Join(dossier, nom))
		if err != nil {
			continue
		}
		j, err := jetonSPDX(string(brut))
		if err != nil {
			return "", "", fmt.Errorf("%s : %w", nom, err)
		}
		offertes[j] = string(brut)
	}

	switch {
	case len(offertes) == 0:
		return "", "", fmt.Errorf("aucun fichier de licence dans %s", dossier)
	case len(offertes) == 1:
		for j, t := range offertes {
			return t, j, nil
		}
	}

	if t, ok := offertes[licenceExercee]; ok {
		return t, licenceExercee, nil
	}

	jetons := make([]string, 0, len(offertes))
	for j := range offertes {
		jetons = append(jetons, j)
	}
	sort.Strings(jetons)
	return "", "", fmt.Errorf("offert sous %s, et aucune n'est %s : choisir laquelle est exercée",
		strings.Join(jetons, " ou "), licenceExercee)
}

// extraitApache et extraitMIT portent les fragments normatifs que la
// reconnaissance exige, sans recopier deux licences entières dans un test.
const (
	extraitApache = "                                 Apache License\n" +
		"                           Version 2.0, January 2004\n"
	extraitMIT = "MIT License\n\nPermission is hereby granted, free of charge, to any person " +
		"obtaining a copy\n\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND\n"
)

// TestLicenceDeExercesTheRepositoryLicence garde le cas qui a fait échouer la
// génération : un module offert sous deux licences au choix.
//
// Le cas sans Apache-2.0 compte autant que celui avec. Retenir la seule restante
// paraîtrait serviable et engagerait la redistribution sur un régime que
// personne n'aurait arbitré — c'est précisément ce que l'erreur refuse.
func TestLicenceDeExercesTheRepositoryLicence(t *testing.T) {
	cas := []struct {
		nom      string
		fichiers map[string]string
		jeton    string
		erreur   string
	}{
		{
			nom:      "double offre",
			fichiers: map[string]string{"LICENSE-APACHE": extraitApache, "LICENSE-MIT": extraitMIT},
			jeton:    "Apache-2.0",
		},
		{
			nom:      "licence unique, même si ce n'est pas celle du dépôt",
			fichiers: map[string]string{"COPYING": extraitMIT},
			jeton:    "MIT",
		},
		{
			nom:      "fichier de licence illisible par la table",
			fichiers: map[string]string{"LICENSE": "Copyright 2026. Tous droits réservés."},
			erreur:   "licence non reconnue",
		},
		{
			nom:      "aucun fichier de licence",
			fichiers: map[string]string{"README.md": "rien ici"},
			erreur:   "aucun fichier de licence",
		},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			dossier := t.TempDir()
			for nom, contenu := range c.fichiers {
				if err := os.WriteFile(filepath.Join(dossier, nom), []byte(contenu), 0o600); err != nil {
					t.Fatal(err)
				}
			}

			_, jeton, err := licenceDe(dossier)
			if c.erreur != "" {
				if err == nil {
					t.Fatalf("jeton %q rendu, attendu une erreur portant %q", jeton, c.erreur)
				}
				if !strings.Contains(err.Error(), c.erreur) {
					t.Errorf("erreur %q, attendu qu'elle porte %q", err, c.erreur)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if jeton != c.jeton {
				t.Errorf("jeton %q, attendu %q", jeton, c.jeton)
			}
		})
	}
}

// TestLicenceDeRefusesToChooseWithoutTheExercisedOne garde le refus de choisir
// quand une double offre ne comprend pas la licence du dépôt.
//
// **La table est substituée, et c'est la seule façon d'atteindre ce cas** :
// elle ne connaît que les licences du binaire, donc deux offres y comprennent
// toujours Apache-2.0 et la branche ne se déclenche jamais en l'état. La laisser
// sans test en ferait du code non gardé, qui s'ouvrirait en silence le jour où
// une troisième licence y entrera — et c'est précisément ce jour-là qu'un
// arbitraire se glisserait dans une notice.
func TestLicenceDeRefusesToChooseWithoutTheExercisedOne(t *testing.T) {
	const extraitISC = "Permission to use, copy, modify, and/or distribute this software\n\n" +
		"THE AUTHOR DISCLAIMS ALL WARRANTIES WITH REGARD TO THIS SOFTWARE\n"

	originale := licencesConnues
	t.Cleanup(func() { licencesConnues = originale })
	licencesConnues = []licenceConnue{
		{"MIT", []string{"Permission is hereby granted, free of charge", `THE SOFTWARE IS PROVIDED "AS IS"`}},
		{"ISC", []string{"Permission to use, copy, modify, and/or distribute", "WITH REGARD TO THIS SOFTWARE"}},
	}

	dossier := t.TempDir()
	for nom, contenu := range map[string]string{"LICENSE-MIT": extraitMIT, "LICENCE": extraitISC} {
		if err := os.WriteFile(filepath.Join(dossier, nom), []byte(contenu), 0o600); err != nil {
			t.Fatal(err)
		}
	}

	_, jeton, err := licenceDe(dossier)
	if err == nil {
		t.Fatalf("jeton %q retenu alors qu'aucune offre n'est %s", jeton, licenceExercee)
	}
	for _, attendu := range []string{"ISC ou MIT", licenceExercee} {
		if !strings.Contains(err.Error(), attendu) {
			t.Errorf("erreur %q, attendu qu'elle porte %q", err, attendu)
		}
	}
}

// TestSPDXTokenRefusesWhatItCannotName vérifie que la reconnaissance échoue
// franchement plutôt que de rendre un jeton par défaut.
//
// C'est la propriété qui donne sa valeur au fichier : un identifiant SPDX écrit
// dans une notice affirme un régime juridique, et l'affirmer de travers est pire
// que ne rien affirmer.
func TestSPDXTokenRefusesWhatItCannotName(t *testing.T) {
	if _, err := jetonSPDX("Copyright 2026. Tous droits réservés."); err == nil {
		t.Error("un texte qui n'est aucune licence connue reçoit un jeton")
	}
	if _, err := jetonSPDX(extraitApache + extraitMIT); err == nil {
		t.Error("un texte reconnu deux fois reçoit un jeton au lieu d'arrêter le contrôle")
	}
}

// TestEveryLinkedModuleIsNamedInTheNotices vérifie que chaque module lié porte
// son identifiant SPDX dans le fichier livré.
//
// Le test voisin compare le fichier à ce que la génération produit, donc il
// rougirait aussi. Celui-ci dit *quoi* est faux : une entrée sans jeton est une
// notice qui ne nomme pas le régime qu'elle impose au redistributeur.
func TestEveryLinkedModuleIsNamedInTheNotices(t *testing.T) {
	brut, err := os.ReadFile(cheminNotices)
	if err != nil {
		t.Fatalf("%v — lancer « make notices »", err)
	}

	modules, err := modulesLies()
	if err != nil {
		t.Fatal(err)
	}
	if len(modules) == 0 {
		t.Fatal("aucun module lié : le contrôle ne dirait rien")
	}

	for _, m := range modules {
		_, jeton, err := licenceDe(m.dossier)
		if err != nil {
			t.Fatalf("%s : %v", m.chemin, err)
		}
		attendu := fmt.Sprintf("%s %s — %s", m.chemin, m.version, jeton)
		if !strings.Contains(string(brut), attendu) {
			t.Errorf("les notices ne portent pas %q", attendu)
		}
	}
}

// jetonSPDX reconnaît une licence à son texte.
//
// Une correspondance et une seule : zéro dit que la table ne connaît pas ce
// texte, plusieurs qu'elle ne discrimine plus, et les deux doivent arrêter le
// contrôle plutôt que rendre le premier jeton venu.
func jetonSPDX(texte string) (string, error) {
	var trouves []string
	for _, l := range licencesConnues {
		reconnu := true
		for _, m := range l.marqueurs {
			if !strings.Contains(texte, m) {
				reconnu = false
				break
			}
		}
		if reconnu {
			trouves = append(trouves, l.jeton)
		}
	}

	switch len(trouves) {
	case 1:
		return trouves[0], nil
	case 0:
		return "", fmt.Errorf("licence non reconnue : ajouter ses marqueurs à licencesConnues")
	default:
		return "", fmt.Errorf("texte reconnu comme %s à la fois : les marqueurs ne discriminent plus",
			strings.Join(trouves, " et "))
	}
}
