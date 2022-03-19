// bingo package is responsible for solving a bingo
package bingo

type Board [][]int

// GetWinnerBoardAndScore returns wich board has won the bingo and with which score.
func GetWinnerBoardAndScore(numbersToDraw []int, boards []Board, wantFirstWinner bool) (winnerBoardIndex int, boardScore int) {
	winnerBoardIndexResult := 0
	boardScoreResult := 0

	boardsInfo := createBoardsInfo(boards)

	for _, x := range numbersToDraw {
		winnerBoardIndex, boardScore = getWinningBoard(x, boards, boardsInfo, wantFirstWinner)
		if winnerBoardIndex != -1 {
			if wantFirstWinner {
				return winnerBoardIndex, boardScore
			}

			winnerBoardIndexResult = winnerBoardIndex
			boardScoreResult = boardScore
		}
	}

	return winnerBoardIndexResult, boardScoreResult
}

// getWinningBoard returns the index of the first winning board and the boardScore
// in case any board wins if wantFirstWinner is true.
// If wantFirstWinner is false returns the same but for the last winner.
// If no winner, returns -1 and -1.
func getWinningBoard(x int, boards []Board, boardsInfo []BoardInfo, wantFirstWinner bool) (int, int) {
	iResult := -1
	boardScoreResult := -1

	for i := range boardsInfo {
		if !boardsInfo[i].boardHasWon {
			boardScore, isBoardWinner := doesBoardWin(x, boards[i], &boardsInfo[i])
			if isBoardWinner {
				if wantFirstWinner {
					return i, boardScore
				}

				iResult = i
				boardScoreResult = boardScore
			}
		}
	}
	return iResult, boardScoreResult
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
					boardInfo.boardHasWon = true
					return calculateBoardScore(x, *boardInfo), true
				}
			}
		}
	}
	return -1, false
}

func calculateBoardScore(lastPickedNumber int, boardInfo BoardInfo) int {
	return lastPickedNumber * boardInfo.SumOfUnmarkedPositions
}

// DoesExistBingoInRowOrColumn returns true if the whole row "row" is true
// or the whole column "col" is true. Otherwise, it returns false.
func doesExistBingoInRowOrColumn(values [][]bool, row, col int) bool {
	return isRowFull(values, row) || isColumnFull(values, col)
}

func isRowFull(values [][]bool, rowIndex int) bool {
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

func isColumnFull(values [][]bool, colIndex int) bool {
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
