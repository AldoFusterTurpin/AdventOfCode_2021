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
					[]bingo.Row{
						bingo.Row{
							[]bingo.Cell{
								{22, false}, {13, false}, {17, false}, {11, false}, {0, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{8, false}, {2, false}, {23, false}, {4, false}, {24, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{21, false}, {9, false}, {14, false}, {16, false}, {7, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{6, false}, {10, false}, {3, false}, {18, false}, {5, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{1, false}, {12, false}, {20, false}, {15, false}, {19, false},
							},
						},
					},
				},
				bingo.Board{
					[]bingo.Row{
						bingo.Row{
							[]bingo.Cell{
								{3, false}, {15, false}, {0, false}, {2, false}, {22, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{9, false}, {18, false}, {13, false}, {17, false}, {5, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{19, false}, {8, false}, {7, false}, {25, false}, {23, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{20, false}, {11, false}, {10, false}, {24, false}, {4, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{14, false}, {21, false}, {16, false}, {12, false}, {6, false},
							},
						},
					},
				},
				bingo.Board{
					[]bingo.Row{
						bingo.Row{
							[]bingo.Cell{
								{14, false}, {21, false}, {17, false}, {24, false}, {4, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{10, false}, {16, false}, {15, false}, {9, false}, {19, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{18, false}, {8, false}, {23, false}, {26, false}, {20, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{22, false}, {11, false}, {13, false}, {6, false}, {5, false},
							},
						},
						bingo.Row{
							[]bingo.Cell{
								{2, false}, {0, false}, {12, false}, {3, false}, {7, false},
							},
						},
					},
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
