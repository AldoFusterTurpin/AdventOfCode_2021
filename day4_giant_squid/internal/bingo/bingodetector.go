package bingo

// IDetector is responsible for detecting if exists bingo in a board
// in a specific row or column
type IDetector interface {
	DoesExistBingoInRowOrColumn(values [][]bool, row, col int) bool
}

// Detector is a struct that implements IDetector
type Detector struct {
}

func NewDetector() Detector {
	return Detector{}
}

// DoesExistBingoInRowOrColumn returns true if the whole row "row" is true
// or the whole column "col" is true. Otherwise, it returns false.
func (d Detector) DoesExistBingoInRowOrColumn(values [][]bool, row, col int) bool {
	return d.isFullRowTrue(values, row) || d.isFullColumnTrue(values, col)
}

func (d Detector) isFullRowTrue(values [][]bool, rowIndex int) bool {
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

func (d Detector) isFullColumnTrue(values [][]bool, colIndex int) bool {
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
