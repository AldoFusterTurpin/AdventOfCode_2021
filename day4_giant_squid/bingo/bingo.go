package bingo

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Board [][]int

// type PosInfo map[int]Position

// type Position struct {
// 	i int
// 	j int
// }

type BoardInfo struct {
	// MarkedPositions        PosInfo
	// TODO: convert MarkedPossitions to a map[number][]Position so I can find in O(1)
	// where is a number and mark it
	MarkedPositions        [][]bool
	SumOfUnmarkedPositions int
}

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

func printBoardsInfo(boardsInfo []BoardInfo) {
	fmt.Println("boardsInfo:")
	for _, boardInfo := range boardsInfo {
		fmt.Println(boardInfo)
	}
}

func GetWinnerBoardAndScore(numbersToDraw []int, boards []Board) (winnerBoardIndex int, boardScore int) {
	boardsInfo := createBoardsInfo(boards)

	for _, x := range numbersToDraw {
		winnerBoardIndex, boardScore = processBoards(x, boards, boardsInfo)
		if winnerBoardIndex != -1 {
			return winnerBoardIndex + 1, boardScore
		}
	}

	return winnerBoardIndex, boardScore
}

func processBoards(x int, boards []Board, boardsInfo []BoardInfo) (int, int) {
	for i := range boardsInfo {
		boardScore, isBoardWinner := processBoard(x, boards[i], &boardsInfo[i])
		if isBoardWinner {
			return i, boardScore
		}
	}
	return -1, -1
}

// func findNumberAndUpdateBoardInfo(x int, board Board, boardInfo BoardInfo) {
// 	for i, row := range board {
// 		for j := range row {
// 			if
// 		}
// 	}
// }

func isRowMarked(boardInfo BoardInfo, rowI int) bool {
	row := boardInfo.MarkedPositions[rowI]

	for _, isCellMarked := range row {
		if !isCellMarked {
			return false
		}
	}
	return true
}

func isColumnMarked(boardInfo BoardInfo, columnI int) bool {
	nRows := len(boardInfo.MarkedPositions)

	for j := 0; j < nRows; j++ {
		isCellMarked := boardInfo.MarkedPositions[j][columnI]
		if !isCellMarked {
			return false
		}
	}

	return true
}

func calculateBoardScore(x int, boardInfo BoardInfo) int {
	fmt.Println("x:", x)
	return x * boardInfo.SumOfUnmarkedPositions
}

func doesExistBingoInPos(boardInfo BoardInfo, row int, column int) bool {
	return isColumnMarked(boardInfo, column) || isRowMarked(boardInfo, row)
}

func processBoard(x int, board Board, boardInfo *BoardInfo) (int, bool) {
	for i, row := range board {
		for j, value := range row {
			if value == x {
				boardInfo.MarkedPositions[i][j] = true
				boardInfo.SumOfUnmarkedPositions -= x

				if doesExistBingoInPos(*boardInfo, j, i) {
					return calculateBoardScore(x, *boardInfo), true
				}
			}
		}
	}
	return -1, false
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
