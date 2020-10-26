package search_matrix_test

import (
	"search_matrix"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	t.Run("Case 1: [0][1] = 31", func(t *testing.T) {
		x, y := search_matrix.SearchPixel(31)
		if x != 0 && y != 1 {
			t.Error("Error calculando la matriz")
		}
	})

	t.Run("Case 2: [0][2] = 32", func(t *testing.T) {
		x, y := search_matrix.SearchPixel(32)
		if x != 0 && y != 2 {
			t.Error("Error calculando la matriz")
		}
	})

	t.Run("Case 3: [4][13] = 196", func(t *testing.T) {
		x, y := search_matrix.SearchPixel(196)
		if x != 4 && y != 13 {
			t.Error("Error calculando la matriz")
		}
	})

	t.Run("Case 4: [15][15] = 240", func(t *testing.T) {
		x, y := search_matrix.SearchPixel(240)
		if x != 15 && y != 15 {
			t.Error("Error calculando la matriz")
		}
	})
}
