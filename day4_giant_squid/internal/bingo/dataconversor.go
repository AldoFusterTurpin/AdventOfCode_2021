package bingo

import (
	"log"
	"strconv"
	"strings"
)

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
