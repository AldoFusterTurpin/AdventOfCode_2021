package bingodetector_test

import (
	"giant_squid/bingodetector"
	"testing"
)

func TestDetector_DoesExistBingoInRowOrColumn(t *testing.T) {
	tt := map[string]struct {
		markedPositions [][]bool
		row, col        int
		expected        bool
	}{
		"should_return_false_when_all_pos_are_false": {
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
		"should_return_true_when_first_row_is_true": {
			[][]bool{
				[]bool{true, true, true, true, true},
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
			},
			0, 0,
			true,
		},
		"should_return_false_when_neither_row_or_column_are_all_true": {
			[][]bool{
				[]bool{false, true, true, true, true},
				[]bool{false, true, true, true, true},
				[]bool{false, true, true, true, true},
				[]bool{false, true, true, true, true},
				[]bool{false, true, true, true, true},
			},
			0, 0,
			false,
		},
		"should_return_true_when_first_col_is_true": {
			[][]bool{
				[]bool{true, false, false, false, false},
				[]bool{true, false, false, false, false},
				[]bool{true, false, false, false, false},
				[]bool{true, false, false, false, false},
				[]bool{true, false, false, false, false},
			},
			2, 0,
			true,
		},
		"should_return_true_when_last_row_is_true": {
			[][]bool{
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{true, true, true, true, true},
			},
			4, 0,
			true,
		},
		"should_return_false_when_row_and_col_are_out_of_bounds": {
			[][]bool{
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{true, true, true, true, true},
			},
			-1, -1,
			false,
		},
		"should_return_true_when_row_is_out_of_bounds_but_bingo_in_column": {
			[][]bool{
				[]bool{false, false, false, true, false},
				[]bool{false, false, false, true, false},
				[]bool{false, false, false, true, false},
				[]bool{false, false, false, true, false},
				[]bool{true, false, true, true, false},
			},
			46, 3,
			true,
		},
		"should_return_false_when_col_is_out_of_bounds_and_NO_bingo_in_row": {
			[][]bool{
				[]bool{false, false, true, false, false},
				[]bool{false, false, true, false, false},
				[]bool{false, false, true, false, false},
				[]bool{false, false, true, false, false},
				[]bool{false, false, true, false, false},
			},
			2, 5,
			false,
		},
		"should_return_true_when_col_is_out_of_bounds_and_is_bingo_in_row": {
			[][]bool{
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
				[]bool{true, true, true, true, true},
				[]bool{false, false, false, false, false},
				[]bool{false, false, false, false, false},
			},
			2, 5,
			true,
		},
	}

	for name, tc := range tt {
		d := bingodetector.Detector{}
		got := d.DoesExistBingoInRowOrColumn(tc.markedPositions, tc.row, tc.col)

		if got != tc.expected {
			t.Fatalf("%s: expected %v, but got: %v", name, tc.expected, got)
		}
	}
}
