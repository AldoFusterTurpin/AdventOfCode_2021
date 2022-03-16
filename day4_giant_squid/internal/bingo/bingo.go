// bingo package is responsible for solving a bingo
package bingo

type Board [][]int

// GetFirstWinnerBoardAndScore returns wich board has won the bingo and with which score.
func GetFirstWinnerBoardAndScore(numbersToDraw []int, boards []Board) (winnerBoardIndex int, boardScore int) {
	boardsInfo := createBoardsInfo(boards)

	for _, x := range numbersToDraw {
		winnerBoardIndex, boardScore = getFirstWinningBoard(x, boards, boardsInfo)
		if winnerBoardIndex != -1 {
			return winnerBoardIndex, boardScore
		}
	}

	return winnerBoardIndex, boardScore
}

func calculateBoardScore(lastPickedNumber int, boardInfo BoardInfo) int {
	return lastPickedNumber * boardInfo.SumOfUnmarkedPositions
}

// DoesExistBingoInRowOrColumn returns true if the whole row "row" is true
// or the whole column "col" is true. Otherwise, it returns false.
func doesExistBingoInRowOrColumn(values [][]bool, row, col int) bool {
	return isFullRowTrue(values, row) || isFullColumnTrue(values, col)
}

func isFullRowTrue(values [][]bool, rowIndex int) bool {
	if rowIndex < 0 || rowIndex >= len(values) {
		return false
	}

	row := values[rowIndex]
	for _, v := range row {
		if !v {
			return false
		}
	}
	return true
}

func isFullColumnTrue(values [][]bool, colIndex int) bool {
	if colIndex < 0 {
		return false
	}

	nRows := len(values)
	for i := 0; i < nRows; i++ {
		colsInRow := len(values[i])
		if colIndex >= colsInRow {
			return false
		}

		if !values[i][colIndex] {
			return false
		}
	}

	return true
}

// getFirstWinningBoard returns the index of the first winning board and the boardScore
// in case any board wins. Otherwise returns -1 and -1
func getFirstWinningBoard(x int, boards []Board, boardsInfo []BoardInfo) (int, int) {
	for i := range boardsInfo {
		boardScore, isBoardWinner := doesBoardWin(x, boards[i], &boardsInfo[i])
		if isBoardWinner {
			return i, boardScore
		}
	}
	return -1, -1
}

// doesBoardWin returns the score of the board "board" and true in case the baord wins.
// Otherwise returns -1 and false.
func doesBoardWin(x int, board Board, boardInfo *BoardInfo) (int, bool) {
	for i, row := range board {
		for j, value := range row {
			if value == x {
				boardInfo.MarkedPositions[i][j] = true
				boardInfo.SumOfUnmarkedPositions -= x

				booleanMatrix := boardInfo.MarkedPositions
				if doesExistBingoInRowOrColumn(booleanMatrix, i, j) {
					return calculateBoardScore(x, *boardInfo), true
				}
			}
		}
	}
	return -1, false
}
