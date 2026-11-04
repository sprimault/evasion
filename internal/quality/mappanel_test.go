// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

package quality

import (
	"testing"

	"github.com/sprimault/evasion/internal/core"
	"github.com/sprimault/evasion/internal/render"
)

// TestMapPanelHoldsItsPixelsPerCell vérifie que les bornes du panneau de carte
// donnent bien l'échelle que docs/architecture.md en tire.
//
// **En pixels de surface, qui est l'unité du panneau.** La carte vit dans la
// planche que le rendu agrandit ensuite d'un facteur entier, donc ces pixels-là
// sont ceux où elle est peinte ; à l'écran, tout est multiplié par le même
// facteur et le rapport ne bouge pas.
//
// Le document ne citait qu'une des deux valeurs, douze, en la présentant comme
// « ce qu'il doit montrer » ; l'autre, huit, apparaissait deux phrases plus loin
// pour justifier le rendu en fond de case. Ce sont les deux bouts du même
// intervalle, et c'est la borne basse qui commande la règle du halo.
func TestMapPanelHoldsItsPixelsPerCell(t *testing.T) {
	cote := 0
	for _, pre := range core.Presets() {
		cote = max(cote, pre.Settings.Size)
	}
	if cote == 0 {
		t.Fatal("aucun préréglage : le contrôle ne dit rien")
	}

	bas := float64(render.MapPanelMin) / float64(cote)
	haut := float64(render.MapPanelMax) / float64(cote)
	t.Logf("panneau de %d à %d px sur un plateau de %d cases : de %.1f à %.1f px par case",
		render.MapPanelMin, render.MapPanelMax, cote, bas, haut)

	// Huit pixels par case est le plancher que le document tient : au-dessous,
	// un pion ne se distingue plus d'une case et la carte cesse de porter la
	// vue d'ensemble qui est sa raison d'être.
	if bas < 8 {
		t.Errorf("au plus étroit, le panneau ne donne que %.1f px par case sur le plus "+
			"grand préréglage, sous les huit dont le fond de case dépend", bas)
	}
	// Treize, et non davantage : au-delà, la carte deviendrait une seconde vue
	// de jeu et prendrait la place de l'isométrique.
	if haut > 13 {
		t.Errorf("au plus large, le panneau donne %.1f px par case : il mange l'isométrique", haut)
	}
	if render.MapPanelRatio <= 0 || render.MapPanelRatio >= 0.5 {
		t.Errorf("le panneau prend %.0f %% de la largeur", 100*render.MapPanelRatio)
	}
}
