// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

package render

import "testing"

// TestSpanKeepsTheTwoToOneRatio vérifie qu'une emprise reste deux fois plus
// large que haute.
//
// C'est l'erreur que Span existe pour empêcher, et elle ne se voit pas : un
// calcul qui confond cases d'écran et cases de plateau donne un résultat
// plausible, seulement faux d'un facteur deux.
func TestSpanKeepsTheTwoToOneRatio(t *testing.T) {
	for _, cases := range []int{1, 9, 17, 41} {
		l, h := Span(cases)
		if l != 2*h {
			t.Errorf("Span(%d) = %d x %d, attendu un rapport 2:1", cases, l, h)
		}
		if l != cases*TileWidth {
			t.Errorf("Span(%d) largeur = %d, attendu %d", cases, l, cases*TileWidth)
		}
	}
}

// TestSpanAgreesWithToScreen vérifie que l'emprise annoncée est bien celle que
// la projection produit.
//
// Les deux fonctions doivent rester d'accord : Span n'est utile que si elle
// décrit la même géométrie que ToScreen, sans quoi dimensionner une fenêtre
// avec l'une et y placer des cases avec l'autre laisserait un écart silencieux.
func TestSpanAgreesWithToScreen(t *testing.T) {
	const cotes = 17

	// Les quatre coins d'un carré de cotes cases, projetés.
	gauche, _ := ToScreen(0, cotes-1)
	droite, _ := ToScreen(cotes-1, 0)
	_, haut := ToScreen(0, 0)
	_, bas := ToScreen(cotes-1, cotes-1)

	l, h := Span(cotes)
	if got := droite - gauche + TileWidth; got != l {
		t.Errorf("largeur projetée %d, Span annonce %d", got, l)
	}
	if got := bas - haut + TileHeight; got != h {
		t.Errorf("hauteur projetée %d, Span annonce %d", got, h)
	}
}

// TestStrokeWidthStaysWithinItsBounds vérifie le plancher et le plafond, en
// pixels de planche.
//
// Le plancher est ce qui empêche un trait de tomber à zéro pixel — une unité
// n'en vaut qu'un demi —, le plafond ce qui l'empêche d'avaler une forme
// minuscule. Les deux se contredisent sur une très petite forme, et c'est le
// plancher qui l'emporte : mieux vaut un trait trop épais qu'un trait absent.
//
// Aucune échelle n'entre dans les cas : l'agrandissement étant entier, ces
// valeurs sont celles de tous les paliers.
func TestStrokeWidthStaysWithinItsBounds(t *testing.T) {
	cas := []struct {
		nom      string
		unites   int
		minForme int
		attendu  int
	}{
		{"liseré sur une tête de fugitif", RimWidth, 14, 1},
		{"contour d'une unité, arrondi au supérieur", 1, 14, 1},
		{"trait large sur une grande forme", 4, 40, 2},
		{"plafond atteint", 8, 24, 2},
		{"plancher l'emporte sur le plafond", RimWidth, 6, 1},
	}

	for _, c := range cas {
		if got := StrokeWidth(c.unites, c.minForme); got != c.attendu {
			t.Errorf("%s : StrokeWidth(%d, %d) = %d, attendu %d",
				c.nom, c.unites, c.minForme, got, c.attendu)
		}
	}
}

// TestStrokeUnitsConvertsBackToTheSamePixels vérifie que les deux unités ne
// divergent pas, par l'aller-retour.
//
// C'est la seule garantie qui compte ici : l'aperçu dessine en géométrie et le
// rendu en pixels de planche. Une épaisseur convertie qui ne reviendrait pas sur
// ses pixels ferait montrer à l'aperçu autre chose que le jeu, et c'est
// précisément ce que l'aperçu existe pour éviter.
//
// L'aller-retour plutôt que la formule recopiée : comparer StrokeUnits au calcul
// qu'elle fait elle-même ne peut pas échouer.
func TestStrokeUnitsConvertsBackToTheSamePixels(t *testing.T) {
	for _, unites := range []int{1, RimWidth, 3, 4, 8} {
		for _, minForme := range []int{6, 14, 24, 40} {
			unitesRendues := StrokeUnits(unites, minForme)
			if retour := unitesRendues * RasterTileWidth / TileWidth; retour != StrokeWidth(unites, minForme) {
				t.Errorf("StrokeUnits(%d, %d) = %d unités, qui reviennent à %d pixels "+
					"au lieu de %d", unites, minForme, unitesRendues, retour,
					StrokeWidth(unites, minForme))
			}
		}
	}
}

// La propriété que les deux bornes existent pour tenir — un pion garde un
// remplissage majoritaire — se vérifie dans internal/quality, sur les formes
// réellement livrées. Elle vivait ici avec un diamètre de quatorze en dur,
// présenté comme la plus petite tête d'un pion livré : c'est celle du fugitif,
// et le test n'exerçait donc que la forme la plus favorable.

// TestRequiredColorsCoverBothOutlines vérifie que le socle porte les contours.
//
// Ce sont eux qui détachent une forme d'un sol dont la luminance varie du
// simple au double et demi ; les oublier ne casse rien au chargement et rend les
// pièces illisibles à l'écran.
func TestRequiredColorsCoverBothOutlines(t *testing.T) {
	presentes := map[string]bool{}
	for _, c := range RequiredColors {
		presentes[c] = true
	}

	for _, attendue := range []string{"backdrop", "marker_outline", "fugitive_detail", "inspector_detail"} {
		if !presentes[attendue] {
			t.Errorf("%s manque au socle obligatoire", attendue)
		}
	}
}

// TestGroundGrainStaysWithinItsAmplitude vérifie que le grain reste dans les
// bornes annoncées.
//
// Un écart qui déborderait ne se verrait pas comme un défaut mais comme une
// case anormalement claire ou sombre, prise pour une zone par le joueur.
func TestGroundGrainStaysWithinItsAmplitude(t *testing.T) {
	for _, graine := range []int64{0, 1, -1, 7, 1 << 40, -(1 << 40)} {
		for colonne := range 41 {
			for ligne := range 41 {
				got := GroundGrain(graine, colonne, ligne)
				if got < -GroundGrainAmplitude || got > GroundGrainAmplitude {
					t.Fatalf("grain %d en (%d, %d) pour la graine %d, hors de ±%d",
						got, colonne, ligne, graine, GroundGrainAmplitude)
				}
			}
		}
	}
}

// TestGroundGrainMatchesItsReference fige le grain de quelques cases.
//
// Rien n'oblige un hachage à rester le même, mais en changer redessine toutes
// les villes déjà jouées : une capture d'écran cesse d'être reproductible depuis
// sa graine. Le test échoue alors, ce qui force à le décider plutôt qu'à le
// subir.
func TestGroundGrainMatchesItsReference(t *testing.T) {
	// Des littéraux, et c'est tout le test. La table se remplissait par des
	// appels à GroundGrain avant d'être comparée à GroundGrain : l'assertion
	// valait f(x) == f(x), et muter le hachage ne la faisait pas rougir.
	//
	// Dix cases et non trois : sur trois, deux valeurs coïncidaient, et un
	// hachage changé avait toutes les chances de retomber dessus. Ce qui décide
	// ici est l'étendue couverte, pas le nombre — graines voisines, graine
	// négative, cases alignées et cases éloignées.
	fige := map[[3]int]int{
		{7, 0, 0}:   -4,
		{7, 1, 0}:   -4,
		{7, 0, 1}:   0,
		{7, 12, 30}: 0,
		{7, 20, 20}: 0,
		{7, 40, 40}: -4,
		{1, 0, 0}:   5,
		{1, 12, 30}: 2,
		{99, 5, 7}:  -1,
		{-3, 8, 8}:  5,
	}

	for cle, attendu := range fige {
		if got := GroundGrain(int64(cle[0]), cle[1], cle[2]); got != attendu {
			t.Errorf("grain en (%d, %d) pour la graine %d : %d, attendu %d",
				cle[1], cle[2], cle[0], got, attendu)
		}
	}
}

// TestGroundGrainFollowsTheSeed vérifie que deux graines ne donnent pas le même
// plateau grainé.
//
// Sans cela le grain serait une propriété de la position seule, et toutes les
// villes porteraient exactement les mêmes taches.
func TestGroundGrainFollowsTheSeed(t *testing.T) {
	var differences int
	for colonne := range 41 {
		for ligne := range 41 {
			if GroundGrain(1, colonne, ligne) != GroundGrain(2, colonne, ligne) {
				differences++
			}
		}
	}

	// Onze valeurs possibles : deux graines coïncident sur environ un onzième
	// des cases, et le seuil laisse largement la place au hasard.
	if differences < 41*41/2 {
		t.Errorf("%d cases diffèrent sur %d, deux graines se ressemblent trop", differences, 41*41)
	}
}

// TestGroundGrainSpreadsAcrossItsRange vérifie que le grain ne se concentre pas
// sur quelques valeurs.
//
// Une fonction de hachage mal mêlée passerait les tests précédents en ne rendant
// que deux ou trois écarts, et le sol se lirait alors comme un damier.
func TestGroundGrainSpreadsAcrossItsRange(t *testing.T) {
	vus := map[int]int{}
	for colonne := range 41 {
		for ligne := range 41 {
			vus[GroundGrain(7, colonne, ligne)]++
		}
	}

	if len(vus) != 2*GroundGrainAmplitude+1 {
		t.Errorf("%d écarts distincts sur %d possibles", len(vus), 2*GroundGrainAmplitude+1)
	}
}
