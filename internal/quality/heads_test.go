// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

package quality

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/sprimault/evasion/internal/loader"
	"github.com/sprimault/evasion/internal/render"
	"github.com/sprimault/evasion/plugins"
)

// TestPieceHeadsOnTheSheet vérifie la taille des têtes de pions en pixels de
// planche, que deux textes citent.
//
// docs/contrat-formes.md et la godoc de StrokeWidth s'en servent pour justifier
// le plafond d'épaisseur : à cette taille, deux pixels de liseré avaleraient une
// tête. L'argument ne vaut que si le chiffre est le bon, et il a déjà annoncé
// « quatre et demi » pour deux têtes qu'aucune géométrie livrée ne produisait.
//
// Les deux tombent sur un entier, et ce n'est pas un hasard : c'est la propriété
// qui a fait retenir 32 pixels par case plutôt que 48, où chaque épaisseur
// tomberait sur un demi-pixel et l'arrondi déciderait seul.
//
// Mesuré sur les formes livrées plutôt que recopié : c'est shapes.toml qui
// décide, et il bougera à l'étape 7 quand le rendu existera pour de bon.
func TestPieceHeadsOnTheSheet(t *testing.T) {
	_, formes, err := loader.Load(plugins.Shipped(), "")
	if err != nil {
		t.Fatalf("chargement du contenu livré : %v", err)
	}

	attendues := map[string]int{"fugitive": 7, "inspector": 6}
	for nom, attendu := range attendues {
		tete := teteDe(t, formes, nom)
		if got := tete * render.RasterTileWidth / render.TileWidth; got != attendu {
			t.Errorf("tête de %s à %d pixels de planche, les textes disent %d",
				nom, got, attendu)
		}
	}
}

// TestPieceCoreStaysMajority vérifie qu'un pion garde un remplissage
// majoritaire, sur la plus petite tête livrée et non sur la plus grande.
//
// C'est la couleur du noyau qui dit à quel camp un pion appartient : un noyau
// minoritaire n'est pas un défaut d'esthétique mais une perte d'information.
//
// Le contrôle vivait dans render avec un diamètre de quatorze en dur, présenté
// comme « la plus petite forme d'un pion livré ». C'est celle du fugitif ; celle
// d'un inspecteur fait dix, et il y en a cinq sur le plateau. Le test passait
// donc en n'exerçant que la forme la plus favorable.
func TestPieceCoreStaysMajority(t *testing.T) {
	_, formes, err := loader.Load(plugins.Shipped(), "")
	if err != nil {
		t.Fatalf("chargement du contenu livré : %v", err)
	}

	pire, pireNom := 1.0, ""
	for nom, forme := range formes.Shapes {
		if forme.Role != render.RolePiece {
			continue
		}
		tete := 0
		for _, trait := range forme.Strokes {
			if trait.Type == render.StrokeCircle {
				tete = max(tete, 2*trait.Radius)
			}
		}
		if tete == 0 {
			continue
		}

		if part := remplissage(tete); part < pire {
			pire, pireNom = part, nom
		}
	}

	if pireNom == "" {
		t.Fatal("aucun pion à tête ronde dans le contenu livré : le contrôle ne dit rien")
	}
	t.Logf("pire remplissage : %s à %.1f %%", pireNom, 100*pire)

	if pire <= 0.5 {
		t.Errorf("le noyau de %s ne fait que %.1f %% du pion", pireNom, 100*pire)
	}
}

// remplissage rend la part du noyau dans la largeur d'un pion, bordure comprise.
//
// Tout en pixels de planche, et sans paramètre de taille : l'agrandissement
// étant entier, chaque pixel devient un bloc et la part ne bouge pas d'un palier
// à l'autre. C'est ce qui permet au contrat de publier deux chiffres au lieu
// d'un tableau par échelle.
func remplissage(diametre int) float64 {
	noyau := diametre * render.RasterTileWidth / render.TileWidth
	bordure := render.StrokeWidth(2, diametre) +
		render.StrokeWidth(render.RimWidth, diametre)
	return float64(noyau) / float64(noyau+2*bordure)
}

// TestFillTableMatchesTheMeasure vérifie que le tableau de remplissage du
// contrat porte ce qu'on mesure sur les formes livrées.
//
// Le contrat en publiait trois valeurs, celles du fugitif seul, sous une phrase
// qui parlait d'« un pion ». Un tableau de chiffres qui n'est comparé à rien
// vieillit sans qu'on le voie — c'est déjà arrivé à celui du liseré.
func TestFillTableMatchesTheMeasure(t *testing.T) {
	_, formes, err := loader.Load(plugins.Shipped(), "")
	if err != nil {
		t.Fatalf("chargement du contenu livré : %v", err)
	}

	contenu, err := os.ReadFile(filepath.Join(racine, "docs", "contrat-formes.md"))
	if err != nil {
		t.Fatal(err)
	}
	const entete = "| Pion | Remplissage |"
	_, apres, trouve := strings.Cut(string(contenu), entete)
	if !trouve {
		t.Fatal("le tableau de remplissage est absent du contrat de formes")
	}
	section, _, _ := strings.Cut(apres, "\n\n")

	// Le libellé du contrat est en français, les noms de formes en anglais :
	// c'est la règle de langue du projet, et la table est ce qui les relie.
	formeDe := map[string]string{"Fugitif": "fugitive", "Inspecteur": "inspector"}

	ligne := regexp.MustCompile(`^\|\s*(\p{L}+)\s*\|\s*(\d+) %\s*\|$`)
	lues := map[string]bool{}
	for _, brute := range strings.Split(section, "\n") {
		trouvaille := ligne.FindStringSubmatch(strings.TrimSpace(brute))
		if trouvaille == nil {
			continue
		}

		nom, present := formeDe[trouvaille[1]]
		if !present {
			t.Errorf("le tableau de remplissage nomme %q, qui n'est pas une forme livrée",
				trouvaille[1])
			continue
		}
		lues[nom] = true

		annonce, _ := strconv.Atoi(trouvaille[2])
		if mesure := int(100*remplissage(teteDe(t, formes, nom)) + 0.5); annonce != mesure {
			t.Errorf("le contrat donne %s à %d %%, la mesure %d %%", nom, annonce, mesure)
		}
	}

	for nom := range formeDe {
		if !lues[formeDe[nom]] {
			t.Errorf("le tableau de remplissage ne dit rien de %s", formeDe[nom])
		}
	}
}

// teteDe rend le diamètre de la tête d'un pion livré.
func teteDe(t *testing.T, formes *render.ShapeSet, nom string) int {
	t.Helper()

	forme, present := formes.Shapes[nom]
	if !present {
		t.Fatalf("la forme %s a disparu du contenu livré", nom)
	}
	tete := 0
	for _, trait := range forme.Strokes {
		if trait.Type == render.StrokeCircle {
			tete = max(tete, 2*trait.Radius)
		}
	}
	if tete == 0 {
		t.Fatalf("%s n'a plus de tête : aucun cercle dans ses traits", nom)
	}
	return tete
}
