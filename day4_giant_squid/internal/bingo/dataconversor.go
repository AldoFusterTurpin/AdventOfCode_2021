package bingo

import (
	"bufio"
	"log"
	"os"
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

func ConvertSliceStringsToSliceInts(slice []string) []int {
	var r []int
	for _, s := range slice {
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}

		r = append(r, x)
	}
	return r
}

func GetRawInput(filPath string) (numbersToDraw []string, boards [][]string) {
	file, err := os.Open(filPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numbersToDraw = strings.Split(scanner.Text(), ",") // TODO convert numbersToDraw (type string) to type []int

	scanner.Scan()
	var board []string

	var t string
	for scanner.Scan() {
		t = scanner.Text()

		if t == "" {

			// this line is needed to be able to read the whole file even if there are
			// extra blank lines at the end of the file
			if len(board) > 0 {
				boards = append(boards, board)
			}
			board = nil
		} else {
			board = append(board, strings.TrimSpace(t))
		}
	}

	//this line is key to be able to read whole file even if there are
	// extra blank lines at the end of the file. Boards holds the value of the
	// last iteration of the for loop
	if len(board) > 0 {
		boards = append(boards, board)
	}

	return numbersToDraw, boards
}
