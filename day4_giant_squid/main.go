package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	randomNumbers := scanner.Text()
	fmt.Println("randomNumbers:")
	fmt.Println(randomNumbers)

	scanner.Scan()
	var boards [][]string
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

	//this line is key to be able to read all file even if there are
	// extra blank lines at the end of the file. Boards holds the value of the
	// last iteration of the for loop
	if len(board) > 0 {
		boards = append(boards, board)
	}

	for i, b := range boards {
		fmt.Println("board", i)
		for j, bb := range b {
			fmt.Println("row", j, ":", bb)
		}
		fmt.Println()
	}
}
