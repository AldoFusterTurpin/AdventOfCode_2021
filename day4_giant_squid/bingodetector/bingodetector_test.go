package bingodetector_test

import (
	"giant_squid/bingodetector"
	"testing"
)

func TestDoesExistBingoInPos(t *testing.T) {
	tt := map[string]struct {
		markedPositions [][]bool
		row, col        int
		expected        bool
	}{
		"simple": {
			[][]bool{
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
			},
			0, 0,
			false,
		},
	}

	for name, tc := range tt {
		got := bingodetector.DoesExistBingoInPos(tc.markedPositions, tc.row, tc.col)

		if got != tc.expected {
			t.Fatalf("%s: expected %v, but got: %v", name, tc.expected, got)
		}
	}
}
