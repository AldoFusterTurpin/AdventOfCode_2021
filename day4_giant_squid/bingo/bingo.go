// bingo package is responsible for solving a bingo
package bingo

import (
	"giant_squid/bingodetector"
	"log"
	"strconv"
	"strings"
)

type Solver struct {
	bingoDetector bingodetector.IDetector
}

func NewSolver(bd bingodetector.IDetector) Solver {
	return Solver{
		bingoDetector: bd,
	}
}

type Board [][]int

//TODO: split in more funcitons (don't want 3 fors in same function)
func ConvertRawInputToBoardsType(rawInput [][]string) []Board {
	var boards []Board

	for _, boardStr := range rawInput {
		board := make([][]int, 0)

		for _, rowStr := range boardStr {
			row := make([]int, 0)
			for _, s := range strings.Split(rowStr, " ") {
				if s != " " && s != "" {
					x, err := strconv.Atoi(s)
					if err != nil {
						log.Fatal(err)
					}
					row = append(row, x)
				}
			}

			board = append(board, row)
		}
		boards = append(boards, board)
	}
	return boards
}

// GetWinnerBoardAndScore returns wich board has won the bigno and with which score.
func (bs Solver) GetWinnerBoardAndScore(numbersToDraw []int, boards []Board) (winnerBoardIndex int, boardScore int) {
	boardsInfo := createBoardsInfo(boards)

	for _, x := range numbersToDraw {
		winnerBoardIndex, boardScore = bs.processBoards(x, boards, boardsInfo)
		if winnerBoardIndex != -1 {
			return winnerBoardIndex, boardScore
		}
	}

	return winnerBoardIndex, boardScore
}

// processBoards returns the index of the winning board and the boardScore
// in case any board wins
func (bs Solver) processBoards(x int, boards []Board, boardsInfo []BoardInfo) (int, int) {
	for i := range boardsInfo {
		boardScore, isBoardWinner := bs.processBoard(x, boards[i], &boardsInfo[i])
		if isBoardWinner {
			return i, boardScore
		}
	}
	return -1, -1
}

// processBoard returns the score of the winning board and true in case any baord wins.
// Otherwise returns -1 and false.
func (bs Solver) processBoard(x int, board Board, boardInfo *BoardInfo) (int, bool) {
	for i, row := range board {
		for j, value := range row {
			if value == x {
				boardInfo.MarkedPositions[i][j] = true
				boardInfo.SumOfUnmarkedPositions -= x

				booleanMatrix := boardInfo.MarkedPositions
				if bs.bingoDetector.DoesExistBingoInRowOrColumn(booleanMatrix, i, j) {
					return bs.calculateBoardScore(x, *boardInfo), true
				}
			}
		}
	}
	return -1, false
}

func (bs Solver) calculateBoardScore(lastPickedNumber int, boardInfo BoardInfo) int {
	return lastPickedNumber * boardInfo.SumOfUnmarkedPositions
}
