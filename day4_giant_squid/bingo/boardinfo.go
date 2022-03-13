package bingo

import "fmt"

type BoardInfo struct {
	// MarkedPositions        PosInfo
	// TODO: convert MarkedPossitions to a map[number][]Position so I can find in O(1)
	// where is a number and mark it
	MarkedPositions        [][]bool
	SumOfUnmarkedPositions int
}

func printBoardsInfo(boardsInfo []BoardInfo) {
	fmt.Println("boardsInfo:")
	for _, boardInfo := range boardsInfo {
		fmt.Println(boardInfo)
	}
}

func createBoardsInfo(boards []Board) []BoardInfo {
	boardsInfo := make([]BoardInfo, len(boards))

	for i, board := range boards {
		boardsInfo[i] = createBoardInfo(board, i)
	}

	return boardsInfo
}

func createBoardInfo(board Board, boardIndex int) BoardInfo {
	nRows := len(board)
	boardInfo := BoardInfo{
		MarkedPositions:        make([][]bool, nRows),
		SumOfUnmarkedPositions: 0,
	}

	for i, row := range board {
		for _, value := range row {
			nCols := len(row)
			boardInfo.MarkedPositions[i] = make([]bool, nCols)
			boardInfo.SumOfUnmarkedPositions += value
		}
	}

	return boardInfo
}
