// bingo package is responsible for solving a bingo
package bingo

import (
	"giant_squid/bingodetector"
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

// GetFirstWinnerBoardAndScore returns wich board has won the bingo and with which score.
func (bs Solver) GetFirstWinnerBoardAndScore(numbersToDraw []int, boards []Board) (winnerBoardIndex int, boardScore int) {
	boardsInfo := createBoardsInfo(boards)

	for _, x := range numbersToDraw {
		winnerBoardIndex, boardScore = bs.getFirstWinningBoard(x, boards, boardsInfo)
		if winnerBoardIndex != -1 {
			return winnerBoardIndex, boardScore
		}
	}

	return winnerBoardIndex, boardScore
}

// getFirstWinningBoard returns the index of the first winning board and the boardScore
// in case any board wins. Otherwise returns -1 and -1
func (bs Solver) getFirstWinningBoard(x int, boards []Board, boardsInfo []BoardInfo) (int, int) {
	for i := range boardsInfo {
		boardScore, isBoardWinner := bs.doesBoardWin(x, boards[i], &boardsInfo[i])
		if isBoardWinner {
			return i, boardScore
		}
	}
	return -1, -1
}

// doesBoardWin returns the score of the board "board" and true in case the baord wins.
// Otherwise returns -1 and false.
func (bs Solver) doesBoardWin(x int, board Board, boardInfo *BoardInfo) (int, bool) {
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
