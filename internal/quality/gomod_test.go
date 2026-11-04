// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

package quality

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

// TestGoModIsTidy vérifie que go.mod et go.sum sont ceux que « go mod tidy »
// produirait.
//
// Rien d'autre ne les regarde : gofmt ignore go.mod, le lint ne lit pas le
// graphe des modules, et govulncheck se tait tant qu'aucun avis ne tombe sur
// une version. Un module laissé non rangé garde donc des exigences que plus
// rien ne justifie, et c'est ainsi que golang.org/x/text y est restée douze
// mineures en arrière, à côté d'une dépendance de test marquée indirecte alors
// qu'un test l'importe.
func TestGoModIsTidy(t *testing.T) {
	var diff, erreurs bytes.Buffer

	commande := exec.Command("go", "mod", "tidy", "-diff")
	commande.Stdout = &diff
	commande.Stderr = &erreurs

	err := commande.Run()
	switch {
	case err == nil:
		return

	// Une panne d'exécution sort du même code qu'un diff, mais n'écrit rien sur
	// la sortie standard. Les confondre ferait accuser go.mod d'un cache de
	// modules absent.
	case diff.Len() == 0:
		t.Fatalf("go mod tidy -diff : %v\n%s", err, strings.TrimSpace(erreurs.String()))

	default:
		t.Errorf("go.mod ou go.sum ne sont pas rangés, « go mod tidy » les change :\n%s",
			strings.TrimRight(diff.String(), "\n"))
	}
}
