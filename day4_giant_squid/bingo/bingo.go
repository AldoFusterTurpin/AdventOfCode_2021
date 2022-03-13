package bingo

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

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

func calculateBoardScore(x int, boardInfo BoardInfo) int {
	fmt.Println("x:", x)
	return x * boardInfo.SumOfUnmarkedPositions
}
