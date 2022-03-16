package main

import (
	"fmt"
	"giant_squid/internal/bingo"
)

func main() {
	fileName := "input_example.txt"
	numbersToDraw, rawBoards := bingo.GetRawInput(fileName)

	numsToDraw := bingo.ConvertSliceStringsToSliceInts(numbersToDraw)
	boards := bingo.ConvertRawInputToBoardsType(rawBoards)

	// bingoDetector := bingo.NewDetector()
	// bingoSolver := bingo.NewFirstWinnerBoardFinder(bingoDetector)

	bingoWinnerIndex, score := bingo.GetWinnerBoardAndScore(numsToDraw, boards)
	fmt.Println("bingoWinnerIndex:", bingoWinnerIndex)
	fmt.Println("score:", score)
}
