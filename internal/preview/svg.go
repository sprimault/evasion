// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

// Package preview rend un jeu de formes en SVG, pour qui met au point un
// plugin d'apparence.
//
// C'est de l'outillage et non le chemin du rendu : en jeu, une forme se
// rastérise une fois au chargement puis se blitte. Le SVG sert à relire un
// plugin sans lancer le jeu, et à voir ce qu'une relecture de catalogue
// accepte.
//
// Le paquet ne décide de rien : il dessine ce que le contrat décrit, avec les
// mêmes règles que le rendu — coefficients de faces, liseré, encadrement des
// épaisseurs. Un aperçu qui s'en écarterait montrerait autre chose que le jeu.
package preview

import (
	"fmt"
	"strings"

	"github.com/sprimault/evasion/internal/render"
)

// trait rend une primitive à l'échelle donnée, liseré compris.
//
// L'ordre compte : liseré, contour, puis remplissage. Un tracé SVG est centré
// sur le chemin et mord donc vers l'intérieur autant que vers l'extérieur ;
// peindre le remplissage en dernier recouvre ce qui a mordu, et il ne subsiste
// que l'épaisseur voulue, à l'extérieur.
func trait(s *strings.Builder, t render.Stroke, x, y, k float64, pal render.Palette) {
	c := pal[t.Color]
	mini := minDimension(t)

	// Nulle sans contour à peindre : la couche du liseré est tracée à
	// epContour + epLisere puis recouverte par le contour, et sans contour rien
	// ne repeignait cette bande — le liseré faisait alors trois unités au lieu
	// de deux, six avec une épaisseur de quatre. Son épaisseur devenait pilotée
	// par le plugin, ce que le contrat exclut.
	// Les épaisseurs viennent du rendu, en unités de contrat, puis suivent k
	// comme la géométrie : l'aperçu montre ce que le jeu produira, bornes
	// comprises, au lieu de les recalculer dans son unité.
	var epContour float64
	if t.Outline != "" {
		epContour = float64(render.StrokeUnits(t.Outlined(), mini)) * k
	}
	epLisere := float64(render.StrokeUnits(render.RimWidth, mini)) * k

	op := ""
	if o := t.Opaque(); o < render.DefaultOpacity {
		op = fmt.Sprintf(` opacity="%.2f"`, float64(o)/100)
	}
	borde := func(couleur string, epaisseur float64) string {
		return fmt.Sprintf(` stroke="%s" stroke-width="%.2f" stroke-linejoin="round"`, couleur, 2*epaisseur)
	}

	switch t.Type {
	case render.StrokePolygon:
		var pts []string
		for _, p := range t.Points {
			pts = append(pts, fmt.Sprintf("%.1f,%.1f", x+float64(p.X)*k, y-float64(p.Y)*k))
		}
		p := strings.Join(pts, " ")
		fmt.Fprintf(s, `<polygon points="%s" fill="none"%s%s/>`, p, borde(render.RimColor, epContour+epLisere), op)
		if t.Outline != "" {
			fmt.Fprintf(s, `<polygon points="%s" fill="none"%s%s/>`, p, borde(pal[t.Outline], epContour), op)
		}
		fmt.Fprintf(s, `<polygon points="%s" fill="%s"%s/>`, p, c, op)

	case render.StrokeCircle:
		cx, cy, r := x+float64(t.Center.X)*k, y-float64(t.Center.Y)*k, float64(t.Radius)*k
		cercle := func(attrs string) {
			fmt.Fprintf(s, `<circle cx="%.1f" cy="%.1f" r="%.1f" fill="none"%s%s/>`, cx, cy, r, attrs, op)
		}
		cercle(borde(render.RimColor, epContour+epLisere))
		if t.Outline != "" {
			cercle(borde(pal[t.Outline], epContour))
		}
		fmt.Fprintf(s, `<circle cx="%.1f" cy="%.1f" r="%.1f" fill="%s"%s/>`, cx, cy, r, c, op)

	case render.StrokeSegment:
		x1, y1 := x+float64(t.From.X)*k, y-float64(t.From.Y)*k
		x2, y2 := x+float64(t.To.X)*k, y-float64(t.To.Y)*k
		corps := float64(t.Thickness) * k
		ligne := func(couleur string, epaisseur float64) {
			fmt.Fprintf(s, `<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="%s" stroke-width="%.2f" stroke-linecap="round"%s/>`,
				x1, y1, x2, y2, couleur, epaisseur, op)
		}
		ligne(render.RimColor, corps+2*(epContour+epLisere))
		if t.Outline != "" {
			ligne(pal[t.Outline], corps+2*epContour)
		}
		ligne(c, corps)

	case render.StrokePrism:
		prisme(s, t, x, y, k, pal)
	}
}

// prisme extrude le losange de la case, trois faces dérivées d'une couleur.
func prisme(s *strings.Builder, t render.Stroke, x, y, k float64, pal render.Palette) {
	dx, dy := render.TileWidth/2*k, render.TileHeight/2*k
	c, h := pal[t.Color], float64(t.Height)*k

	fmt.Fprintf(s,
		`<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="%s"/>`+
			`<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="%s"/>`+
			`<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="%s"/>`,
		x-dx, y-h, x, y-dy-h, x+dx, y-h, x, y+dy-h, render.Lit(c, render.FaceTop),
		x+dx, y-h, x, y+dy-h, x, y+dy, x+dx, y, render.Lit(c, render.FaceRight),
		x-dx, y-h, x, y+dy-h, x, y+dy, x-dx, y, render.Lit(c, render.FaceLeft))
}

// losange trace le sol d'une case, décalé par son grain.
func losange(s *strings.Builder, x, y, k float64, couleur string, grain int) {
	dx, dy := render.TileWidth/2*k, render.TileHeight/2*k
	fmt.Fprintf(s, `<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="%s"/>`,
		x-dx, y, x, y-dy, x+dx, y, x, y+dy, render.ShiftLuminance(couleur, grain))
}

// minDimension renvoie la plus petite dimension d'un trait, en unités de
// contrat. C'est elle que le plafond d'épaisseur ne doit pas laisser avaler.
func minDimension(t render.Stroke) int {
	switch t.Type {
	case render.StrokeCircle:
		return 2 * t.Radius
	case render.StrokeSegment:
		return t.Thickness
	case render.StrokePolygon:
		if len(t.Points) == 0 {
			return 1
		}
		xmin, xmax := t.Points[0].X, t.Points[0].X
		ymin, ymax := t.Points[0].Y, t.Points[0].Y
		for _, p := range t.Points[1:] {
			xmin, xmax = min(xmin, p.X), max(xmax, p.X)
			ymin, ymax = min(ymin, p.Y), max(ymax, p.Y)
		}
		return min(xmax-xmin, ymax-ymin)
	}
	return 1
}

// etiquette écrit un libellé lisible sur un fond quelconque.
func etiquette(s *strings.Builder, x, y float64, texte, couleur string, taille int, ancrage string) {
	fmt.Fprintf(s, `<text x="%.1f" y="%.1f" font-family="sans-serif" font-size="%d" text-anchor="%s" fill="%s">%s</text>`,
		x, y, taille, ancrage, couleur, echapper(texte))
}

// echapper rend un texte sûr dans un document XML.
//
// Un nom de forme vient d'un fichier écrit par un tiers : le passer tel quel
// produirait un SVG invalide au premier chevron, et l'aperçu servirait
// justement à comprendre pourquoi.
func echapper(s string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", `"`, "&quot;").Replace(s)
}
