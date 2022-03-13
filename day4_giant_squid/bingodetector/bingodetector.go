// bingodetector is responsible for detecting if exists bingo in a board
// in a specific row or column
package bingodetector

// DoesExistBingoInRowOrColumn returns true if the whole row "row" is true
// or the whole column "col" is true. Otherwise, it returns false.
func DoesExistBingoInRowOrColumn(values [][]bool, row, col int) bool {
	if row >= 0 && isRowMarked(values, row) {
		return true
	}

	if col >= 0 && isColumnMarked(values, col) {
		return true
	}

	return false
}

func isRowMarked(values [][]bool, rowIndex int) bool {
	if rowIndex >= len(values) {
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

func isColumnMarked(values [][]bool, colIndex int) bool {
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
