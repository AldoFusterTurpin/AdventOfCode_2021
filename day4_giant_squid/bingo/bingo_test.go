package bingo_test

import (
	"giant_squid/bingo"
	"testing"
)

func TestGetWinnerBoardAndScore(t *testing.T) {
	tests := map[string]struct {
		inputBoards              []bingo.Board
		numbersToDraw            []int
		expectedWinnerBingoIndex int
		expectedScore            int
	}{
		"simple": {
			inputBoards: []bingo.Board{
				bingo.Board{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				bingo.Board{
					{3, 15, 0, 2, 22},
					{9, 18, 13, 17, 5},
					{19, 8, 7, 25, 23},
					{20, 11, 10, 24, 4},
					{14, 21, 16, 12, 6},
				},
				bingo.Board{
					{14, 21, 17, 24, 4},
					{10, 16, 15, 9, 19},
					{18, 8, 23, 26, 20},
					{22, 11, 13, 6, 5},
					{2, 0, 12, 3, 7},
				},
			},
			expectedWinnerBingoIndex: 2,
			expectedScore:            4512,
		},
	}

	for name, tc := range tests {
		bingoWinnerIndex, score := bingo.GetWinnerBoardAndScore(tc.numbersToDraw, tc.inputBoards)
		if bingoWinnerIndex != tc.expectedWinnerBingoIndex {
			t.Fatalf("%s: expected %v, but got: %v", name, tc.expectedWinnerBingoIndex, bingoWinnerIndex)
		}
		if score != tc.expectedScore {
			t.Fatalf("%s: expected %v, but got: %v", name, tc.expectedScore, score)
		}
	}
}
