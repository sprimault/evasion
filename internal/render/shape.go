// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

// Package render porte la projection isométrique et le vocabulaire géométrique
// des plugins d'apparence.
//
// C'est le seul paquet qui convertit entre coordonnées de plateau et
// coordonnées d'écran. Le noyau ne connaît que colonne et ligne ; si une
// coordonnée d'écran apparaît ailleurs, c'est un défaut.
package render

import (
	"fmt"
)

// Dimensions du losange, en unités du contrat de formes. Non modifiables : le
// rapport 2:1 est la projection, pas un réglage.
const (
	TileWidth  = 64
	TileHeight = 32
)

// ShapesVersion est la version du contrat que ce binaire sait lire. Un plugin
// écrit contre une version inconnue est refusé plutôt que lu de travers.
const ShapesVersion = 4

// StrokeType énumère les quatre primitives. Le vocabulaire est volontairement
// pauvre : tout ce qui est livré avec le jeu s'y exprime, et une primitive de
// plus entre dans le contrat public sans pouvoir en ressortir.
type StrokeType string

// Les quatre primitives. StrokePrism est réservé au rôle bâtiment et ne porte
// qu'une hauteur : son emprise est le losange de la case, jamais déclarée.
const (
	StrokePolygon StrokeType = "polygon"
	StrokeCircle  StrokeType = "circle"
	StrokeSegment StrokeType = "segment"
	StrokePrism   StrokeType = "prism"
)

// Point est une coordonnée dans le repère d'une forme : origine au point
// d'ancrage au sol, y vers le haut.
type Point struct {
	X int
	Y int
}

// UnmarshalTOML lit un point écrit en paire, « [-6, 0] ».
//
// Le contrat le veut ainsi : une liste de sommets s'écrit alors
// « [[-7, 0], [7, 0], [5, 22]] », lisible d'un coup d'œil, là où une suite de
// tables « {x = -7, y = 0} » noierait la géométrie sous la syntaxe.
func (p *Point) UnmarshalTOML(v any) error {
	paire, ok := v.([]any)
	if !ok || len(paire) != 2 {
		return fmt.Errorf("point: attendu une paire [x, y], reçu %v", v)
	}

	for i, brut := range paire {
		n, ok := brut.(int64)
		if !ok {
			return fmt.Errorf("point: %v n'est pas un entier", brut)
		}
		if i == 0 {
			p.X = int(n)
		} else {
			p.Y = int(n)
		}
	}
	return nil
}

// Stroke est une primitive paramétrée. Les champs inutiles au type restent à
// zéro, comme pour core.Effect et pour la même raison : un enregistrement plat
// se lit et se valide sans hiérarchie de types.
type Stroke struct {
	Type      StrokeType `toml:"type"`
	Points    []Point    `toml:"points"`
	Center    Point      `toml:"center"`
	Radius    int        `toml:"radius"`
	From      Point      `toml:"from"`
	To        Point      `toml:"to"`
	Thickness int        `toml:"thickness"`
	Height    int        `toml:"height"`
	Color     string     `toml:"color"`
	Outline   string     `toml:"outline"`

	// OutlineThickness est un pointeur pour séparer « absent » de « zéro ».
	// Sur un entier, les deux se confondaient : le chargeur acceptait donc un
	// zéro écrit à la main, que le schéma refuse depuis toujours et que son
	// propre message annonçait hors bornes. Absent, l'épaisseur vaut
	// DefaultOutline ; écrite, elle doit tenir dans les bornes du contrat.
	OutlineThickness *int `toml:"outline_thickness"`

	// Opacity est un pointeur pour la même raison qu'OutlineThickness : son
	// défaut vaut 100 et sa borne basse zéro, donc un entier confondrait « je
	// n'en déclare pas » avec « je la veux nulle ». L'aperçu traitait les deux
	// comme opaques, et une forme déclarée transparente s'affichait pleine.
	Opacity *int `toml:"opacity"`
}

// Les défauts d'un trait qui ne les déclare pas.
const (
	DefaultOutline = 1
	DefaultOpacity = 100
)

// Outlined rend l'épaisseur de contour effective d'un trait.
//
// Point unique parce que le défaut se réapplique partout où un contour se
// dessine, et qu'un oubli y donnerait un trait d'épaisseur nulle plutôt que le
// contour attendu.
func (s Stroke) Outlined() int {
	if s.OutlineThickness == nil {
		return DefaultOutline
	}
	return *s.OutlineThickness
}

// Opaque rend l'opacité effective d'un trait, de 0 à 100.
func (s Stroke) Opaque() int {
	if s.Opacity == nil {
		return DefaultOpacity
	}
	return *s.Opacity
}

// Role détermine ce qu'un plugin a le droit de redéfinir.
//
// Il n'existe pas de rôle « sol » : le losange est tracé par le moteur et ne se
// redessine pas. Rue et zones sont des noms de couleurs, pas des formes.
//
// La raison est mécanique. Un losange identique partout, c'est une géométrie
// construite une fois pour toutes les cases, un tri en profondeur qui reste
// colonne + ligne, et un ToBoard qui reste une formule fermée. Une tuile de
// forme libre imposerait une boîte englobante par case, un test de survol forme
// par forme, et invaliderait le tri par ordre de peintre.
type Role string

// Les trois rôles. Le sol n'en est pas un : les cinq noms de Grounds sont des
// couleurs, et le losange est tracé par le moteur.
const (
	RolePiece    Role = "piece"
	RoleBuilding Role = "building"
	RoleMarker   Role = "marker"
)

// Template borne ce qu'une forme a le droit d'occuper.
//
// Vérifié au chargement et pas seulement à la publication : une forme qui
// déborde masque les cases derrière elle, ce qui est un avantage de jeu déguisé
// en habillage. Un plugin local est donc soumis au même contrôle qu'un plugin
// catalogué.
//
// Les bornes portent sur le plan du rôle — celui du sol pour une forme plate et
// pour l'emprise d'un prisme, le plan vertical pour un pion. MaxHeight est
// comptée à part parce qu'un prisme mêle les deux : son emprise est au sol, son
// élévation ne l'est pas.
type Template struct {
	Plane      Plane
	XMin, XMax int
	YMin, YMax int
	MaxStrokes int
	MaxHeight  int
}

// Plane distingue les deux repères. Les confondre est l'erreur classique du
// contrat : un losange au sol occupe y de -16 à 16, un pion s'élève de 0 à 40,
// et les deux plages n'ont rien à voir.
type Plane string

// Les deux repères.
const (
	PlaneGround   Plane = "ground"
	PlaneVertical Plane = "vertical"
)

// Templates borne ce que chaque rôle a le droit d'occuper.
//
// Exportée pour que la table des gabarits de docs/contrat-formes.md se compare à
// elle : le document a donné le marqueur pour vertical alors qu'il est au sol,
// et un contrôle qui lit les deux l'aurait dit.
var Templates = map[Role]Template{
	RolePiece:  {PlaneVertical, -24, 24, 0, 40, 24, 0},
	RoleMarker: {PlaneGround, -24, 24, -12, 12, 8, 0},
	// Un bâtiment ne déclare qu'une hauteur et une couleur : son emprise est le
	// losange de la case. Une emprise libre permettrait de déborder sur les
	// voisines et de masquer ce que l'adversaire doit voir.
	RoleBuilding: {PlaneVertical, 0, 0, 0, 0, 1, 24},
}

// RoleOf impose son rôle à un nom de forme que le jeu consomme, et dit faux
// pour un nom libre.
//
// Le gabarit se choisit sur le rôle déclaré, et le rendu va chercher la forme
// par son nom : sans cette table, une forme « building » déclarée « piece »
// recevait le gabarit d'un pion — emprise libre au lieu du losange imposé — et
// couvrait les cases bâties de rectangles. Le contrat écarte ce cas nommément,
// au motif qu'une emprise libre permettrait « de masquer ce que l'adversaire
// doit voir » ; c'est le nom, et non le rôle, qui décidait de ce qui serait
// dessiné.
//
// Un nom absent est libre : un plugin qui ajoute ses propres formes choisit leur
// rôle, et rien ne va les chercher sous un nom convenu.
func RoleOf(nom string) (Role, bool) {
	// Les cinq noms exactement, et non le préfixe : « inspector_ » fermait tout
	// son espace, si bien qu'un plugin ne pouvait plus déclarer « inspector_halo »
	// en marqueur — refusé au chargement, au nom d'une règle dont le contrat dit
	// deux fois qu'elle ne s'applique pas à lui.
	for i := 1; i <= InspectorOverrides; i++ {
		if nom == fmt.Sprintf("inspector_%d", i) {
			return RolePiece, true
		}
	}
	role, impose := ShapeRoles[nom]
	return role, impose
}

// InspectorOverrides est le nombre de surcharges par pion que le §6 du contrat
// nomme, « inspector_1 » à « inspector_5 » — une par inspecteur.
const InspectorOverrides = 5

// ShapeRoles porte la table de docs/contrat-formes.md §6, les surcharges par
// pion en moins : « inspector_1 » à « inspector_5 » suivent un motif, que RoleOf
// traite à part.
var ShapeRoles = map[string]Role{
	"building":      RoleBuilding,
	"fugitive":      RolePiece,
	"inspector":     RolePiece,
	"trail":         RoleMarker,
	"roadblock":     RoleMarker,
	"cell_visible":  RoleMarker,
	"cell_playable": RoleMarker,
}

// Shape est un dessin nommé, plus ses variantes d'état facultatives.
type Shape struct {
	// La clé de la table porte le nom : « [shape.fugitive] » nomme la forme,
	// et le champ le reçoit au chargement.
	Name string `toml:"-"`

	Role Role `toml:"role"`

	// Les tags sont explicites et non déduits du nom du champ. La déduction
	// est insensible à la casse mais pas à la langue : elle a laissé passer
	// des formes vides quand « trait » est devenu « stroke », sans que rien
	// ne le signale — un TOML dont on ignore une table se décode sans erreur.
	Strokes []Stroke `toml:"stroke"`

	// Les deux variantes d'état, facultatives. Nommées et non rassemblées dans
	// une table : le moteur n'en produit que deux, et une map aurait accepté
	// un état inventé pour l'ignorer ensuite en silence — l'auteur n'aurait
	// appris son erreur qu'en ne la voyant pas à l'écran.
	//
	// Les trois voix du contrat divergeaient ici : le schéma déclarait ces deux
	// propriétés, le document les montrait, et le Go décodait une table sous la
	// clé « variant ». Aucune n'acceptait ce que les deux autres écrivaient.
	Highlighted []Stroke `toml:"highlighted"`
	OutOfSight  []Stroke `toml:"out_of_sight"`
}

// Variants rend les variantes d'état déclarées, indexées par leur nom.
//
// Un accesseur plutôt que le champ : la validation les parcourt toutes de la
// même façon, et l'ordre est fixe pour que deux chargements du même fichier
// signalent le même manquement en premier.
func (s Shape) Variants() map[string][]Stroke {
	variantes := map[string][]Stroke{}
	if len(s.Highlighted) > 0 {
		variantes["highlighted"] = s.Highlighted
	}
	if len(s.OutOfSight) > 0 {
		variantes["out_of_sight"] = s.OutOfSight
	}
	return variantes
}

// RimColor est le liseré que le moteur pose sous le contour des pions et des
// marqueurs. Un plugin n'a rien à en déclarer, et ne peut pas le retirer.
//
// Un contour seul ne suffit pas, et c'est contre-intuitif. Il tient contre le
// sol, qui est clair, mais un pion se dessine par-dessus les cubes situés
// devant lui : sa moitié supérieure est en permanence sur du bâti sombre, où un
// contour sombre ne se voit plus — 1,07 sur la face gauche. L'inverse est vrai
// d'un contour clair, qui meurt sur la rue. Aucune couleur unique ne couvre une
// plage de fonds qui va de 17 à 230 en luminance.
//
// Les deux ensemble la couvrent : quel que soit le fond, l'un des deux traits
// tranche, et le pire cas remonte de 1,07 à 3,81. Ce pire cas est un barrage, et
// non le bâti auquel on penserait : un pion se dessine par-dessus le marqueur de
// sa case, ce qui fait des deux marqueurs des fonds à part entière.
//
// Valeur fixe et non nom de palette : c'est de l'éclairage, au même titre que
// les coefficients de faces d'un prisme. Une palette qui pourrait la déplacer
// pourrait rendre les pions illisibles, ce que le liseré existe pour empêcher.
const RimColor = "#e8e2d4"

// RimWidth est l'épaisseur du liseré, en unités de contrat. Comme toute
// épaisseur de trait, elle est encadrée par StrokeWidth.
const RimWidth = 2

// Le panneau de la carte à plat, en pixels de surface.
//
// L'unité compte : la carte est peinte dans la planche que le rendu agrandit
// ensuite d'un facteur entier, donc ces bornes sont celles où elle est peinte.
// À l'écran, tout est multiplié par le même facteur et le rapport ne bouge pas.
//
// Une part de la largeur, bornée des deux côtés : trop étroit, le plus grand
// préréglage n'a plus assez de pixels par case pour qu'on y suive un itinéraire ;
// trop large, il mange l'isométrique, qui est la vue de jeu.
//
// Les bornes se lisent en pixels par case sur le plus grand préréglage — de huit
// à treize. C'est le huit qui commande la règle du fond de case : à cette
// échelle, un halo autour d'un pion déborderait sur les huit voisines.
//
// 328 et non 320, qui était le chiffre rond posé d'abord : il donnait 7,8 pixels
// par case sur un plateau de 41, soit moins que le huit dont l'argument dépend.
// Un contrôle le mesure, et suivra le jour où le plus grand préréglage bougera.
const (
	MapPanelRatio = 0.26
	MapPanelMin   = 328
	MapPanelMax   = 520
)

// Les trois faces d'un prisme, dans l'ordre où PrismFaces les rend.
const (
	FaceTop = iota
	FaceRight
	FaceLeft
)

// PrismFaces porte les coefficients d'éclairage des trois faces d'un prisme.
//
// Ici et non dans le paquet qui dessine : docs/contrat-formes.md §5 les publie
// comme une règle du moteur, en disant que c'est « exactement le genre de détail
// que deux implémentations règlent différemment sans que personne ne s'en
// aperçoive avant de comparer deux captures ». L'aperçu en est une, le rendu de
// l'étape 7 sera la seconde, et elles ne peuvent les partager que depuis ici —
// c'est preview qui importe render, jamais l'inverse.
//
// Le coefficient s'applique aux trois canaux, ce qui préserve la teinte.
var PrismFaces = [3]float64{1.50, 1.14, 0.72}

// Lit applique le coefficient d'une face à une couleur.
//
// Le débordement se borne à 255 plutôt que de reboucler : un canal saturé reste
// saturé, là qu'un modulo rendrait un dessus de bâtiment plus sombre que ses
// côtés dès que la palette monte.
func Lit(hexa string, face int) string {
	var r, v, b int
	if _, err := fmt.Sscanf(hexa, "#%02x%02x%02x", &r, &v, &b); err != nil {
		return "#000000"
	}
	borne := func(n int) int { return min(int(float64(n)*PrismFaces[face]), 255) }
	return fmt.Sprintf("#%02x%02x%02x", borne(r), borne(v), borne(b))
}

// MinStrokePixels et MaxStrokeRatio encadrent l'épaisseur de tout trait de
// contour, celui d'un plugin comme le liseré du moteur.
//
// Les deux bornes traitent le même défaut par ses deux bouts, et l'une sans
// l'autre ne fait que le déplacer. Sans plancher, une épaisseur déclarée fine
// tombe à zéro pixel de planche — une unité n'en vaut qu'un demi — et le trait
// disparaît au lieu de séparer. Sans plafond, une épaisseur forte finit par
// occuper la forme entière : les deux têtes livrées font 7 et 6 pixels, et deux
// pixels de liseré n'en détacheraient plus rien, ils avaleraient la couleur qui
// est le seul signal d'appartenance à un camp.
//
// Le plancher est en pixels de **planche**, et c'est strictement plus fort qu'un
// pixel d'écran : l'agrandissement étant entier, un pixel de planche en vaut
// autant que le facteur.
//
// Le plafond porte sur la plus petite dimension du trait et non sur la case ni
// sur la forme entière : c'est ce trait-là que le contour doit laisser voir, et
// un trait long et fin serait avalé si l'on bornait par sa plus grande.
//
// Chaque trait est encadré pour lui-même, jamais leur somme : borner le total
// reviendrait à écraser le liseré dès que le plafond mord, alors que c'est lui
// qui porte le contraste sur le bâti.
const (
	MinStrokePixels = 1
	MaxStrokeRatio  = 6
)

// La case rastérisée : 32 pixels de large, 16 de haut, pour les 64 et 32 unités
// que le contrat fige.
//
// Il n'y a pas d'échelle continue. Le rendu peint une planche à cette taille,
// puis l'agrandit vers la fenêtre par un facteur **entier** : les pixels par
// case sont la suite 32, 64, 96, jamais un intervalle. Ce qui varie avec la
// fenêtre est le nombre de cases visibles, pas le grain.
//
// 32 et non 64, qui serait exact aussi : le champ du pion sélectionné fait
// 2r+1 = 17 cases, soit Span(17) à tenir dans l'aire de monde. À 64, il
// demanderait 1088 pixels, donc une fenêtre plancher d'au moins 1416 avec le
// panneau, et le premier agrandissement n'arriverait qu'au-delà de 2832.
//
// 32 et non 48 : c'est le seul palier où toute épaisseur du contrat tombe
// juste. Une unité valant un demi-pixel, un contour de 1 unité atteint le
// plancher d'un pixel, le liseré de 2 en fait un, et les remplissages que le
// contrat publie se retrouvent sans règle d'arrondi — tête du fugitif 7 pixels
// sur 11, celle de l'inspecteur 6 sur 10. À 48, chaque épaisseur tombe sur un
// demi-pixel et l'arrondi déciderait seul.
const (
	RasterTileWidth  = TileWidth / 2
	RasterTileHeight = TileHeight / 2
)

// Le plancher de la surface, en points, sous lequel la fenêtre ne descend pas.
//
// En points et non en pixels : la fenêtre s'ouvre à cette taille quelle que soit
// la densité de l'écran, donc le jeu occupe la même surface visible partout et
// l'agrandissement suit le facteur du moniteur — un sur un écran ordinaire, deux
// à 150 %. En dessous de ce plancher, l'image est rognée et centrée, jamais
// réduite.
//
// Ce qu'il doit contenir est l'aire de monde, panneau déduit : Span(17) en
// pixels rastérisés, soit 544 sur 272, contre 632 sur 540 disponibles une fois
// retirées les 328 du panneau. C'est la largeur qui commande, et il reste 88
// pixels de marge.
//
// La garantie en découle et se simplifie : la surface valant toujours au moins
// ce plancher, un champ qui tient ici tient à toute taille de fenêtre. Un seul
// contrôle au montage remplace donc tout calcul de plafond par fenêtre.
const (
	MinSurfaceWidth  = 960
	MinSurfaceHeight = 540
)

// StrokeWidth donne l'épaisseur d'un trait en pixels de planche.
//
// unites est l'épaisseur déclarée, minForme la plus petite dimension de la forme,
// les deux en unités de contrat. Aucune échelle n'entre ici : l'agrandissement
// étant entier, une épaisseur juste sur la planche l'est à tous les facteurs, et
// le contrat se vérifie donc une fois pour toutes à cette taille.
//
// Le calcul reste entier de bout en bout. Un flottant dans le chemin de rendu
// rendrait l'image dépendante d'un arrondi que rien ne fixe, et l'arrondi est
// ici écrit : au supérieur pour le plancher, par troncature pour le plafond. Un
// liseré qui disparaît coûte plus cher qu'un liseré d'un pixel de trop.
func StrokeWidth(unites, minForme int) int {
	e := (unites*RasterTileWidth + TileWidth - 1) / TileWidth
	if plafond := minForme * RasterTileWidth / (TileWidth * MaxStrokeRatio); e > plafond {
		e = plafond
	}
	return max(e, MinStrokePixels)
}

// StrokeUnits donne la même épaisseur, ramenée en unités de contrat.
//
// Pour ce qui dessine en géométrie déclarative plutôt qu'en pixels — l'aperçu
// SVG. Il montre ainsi l'épaisseur que le rendu produira, bornes comprises, au
// lieu de la recalculer dans son unité et de dériver de lui sans que rien ne le
// signale.
func StrokeUnits(unites, minForme int) int {
	return StrokeWidth(unites, minForme) * TileWidth / RasterTileWidth
}

// GroundGrainAmplitude est l'écart appliqué au sol, en niveaux de luminance sur
// 255.
//
// Absolu et non proportionnel, ce que le contrat écarte nommément : à trois
// pour cent, le grain vaudrait six niveaux sur la rue et moins de trois sur une
// zone fermée, donc il disparaîtrait là où les sols sont les plus serrés. C'est
// aussi l'unité dans laquelle validerEcartDesSols le lit pour poser son seuil,
// et cette godoc a annoncé des pourcents assez longtemps pour que l'aperçu les
// applique.
//
// Un plateau de couleurs pleines est plat à l'œil sur seize cents cases. Le
// grain est dérivé de la position et de la graine, donc stable : le retirer au
// sort à l'affichage produirait un scintillement au défilement, pire que l'uni.
//
// Luminosité seule, jamais la teinte, et le sol seulement — pions et bâtiments
// gardent leur couleur exacte, qui doit rester identifiable d'un coup d'œil.
// Coupé en vue à plat, où l'uniformité aide à lire les surcouches.
const GroundGrainAmplitude = 5

// GroundGrain renvoie l'écart de luminosité d'une case, dans
// [-GroundGrainAmplitude, +GroundGrainAmplitude].
//
// Ce n'est pas une primitive du contrat et aucun plugin n'a à en tenir compte :
// tous en bénéficient, y compris ceux qui ne changent que la palette.
// Haché plutôt que calculé au fil de l'eau : une formule arithmétique simple
// laisse des motifs visibles, des diagonales ou des bandes que l'œil retrouve
// aussitôt sur seize cents cases. FNV-1a est ce que le générateur du noyau
// emploie déjà pour mêler un nom de flux à la graine.
func GroundGrain(graine int64, colonne, ligne int) int {
	const (
		base    uint64 = 14695981039346656037
		premier uint64 = 1099511628211
	)

	// Les conversions cherchent le motif de bits et non la valeur : une graine
	// ou une coordonnée négative doit donner un grain comme un autre. Même
	// raison qu'au générateur du noyau, qui mêle sa graine de la même façon.
	graines := [3]uint64{
		uint64(graine),         // #nosec G115
		uint64(int64(colonne)), // #nosec G115
		uint64(int64(ligne)),   // #nosec G115
	}

	h := base
	for _, v := range graines {
		for i := range 8 {
			h ^= (v >> (8 * i)) & 0xff
			h *= premier
		}
	}

	return int(h%(2*GroundGrainAmplitude+1)) - GroundGrainAmplitude
}

// ShiftLuminance décale une couleur d'un nombre de niveaux, en bornant chaque
// canal. Une valeur mal formée rend du noir plutôt que d'échouer.
//
// Le même décalage sur les trois canaux, donc absolu, et c'est ce qui compte :
// un facteur déplace d'autant moins que la couleur est sombre, si bien qu'il
// resserre les sols là où ils le sont déjà et peut leur faire échanger leur
// rang. C'est ce qui arrivait à la rue et au lieu actif sur la palette livrée.
//
// Ici et non chez l'appelant : le rendu et l'aperçu doivent poser le même
// grain, et deux implémentations d'un geste de trois lignes divergent au
// premier qui le réécrit de mémoire.
func ShiftLuminance(hexa string, niveaux int) string {
	var r, v, b int
	if _, err := fmt.Sscanf(hexa, "#%02x%02x%02x", &r, &v, &b); err != nil {
		return "#000000"
	}
	borne := func(c int) int { return min(max(c+niveaux, 0), 255) }
	return fmt.Sprintf("#%02x%02x%02x", borne(r), borne(v), borne(b))
}

// ToScreen projette une case du plateau en coordonnées d'écran.
func ToScreen(colonne, ligne int) (x, y int) {
	return (colonne - ligne) * TileWidth / 2, (colonne + ligne) * TileHeight / 2
}

// Span donne l'emprise projetée d'un carré de cases, en unités du contrat.
//
// À écrire partout où l'on veut savoir quelle place occupe un nombre de cases :
// dimensionner une fenêtre, tenir un champ de vision, comparer une forme au
// terrain. Le rapport 2:1 fait qu'un carré de n cases est deux fois plus large
// que haut, et l'oublier donne un résultat plausible et faux — le facteur écrit
// à la main s'est trompé trois fois de suite en écrivant ce paquet.
func Span(cases int) (largeur, hauteur int) {
	return cases * TileWidth, cases * TileHeight
}

// ToBoard est la transformation inverse, pour le clic.
//
// L'arrondi doit tomber sur la case que l'utilisateur croit viser, y compris
// près des arêtes du losange où le calcul naïf se trompe d'une case.
func ToBoard(x, y int) (colonne, ligne int) {
	// à implémenter : étape 7
	return 0, 0
}
