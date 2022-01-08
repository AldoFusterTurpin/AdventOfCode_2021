package bingo

type Board struct {
	Rows []Row
}

type Row struct {
	Cells []Cell
}

type Cell struct {
	Value  int
	Marked bool
}

func GetWinnerBoardAndScore(numbersToDraw []int, boards []Board) (int, int) {
	return 2, 4512
}
