// bingo package is responsible for solving a bingo
package bingo

type Board [][]int

func calculateBoardScore(lastPickedNumber int, boardInfo BoardInfo) int {
	return lastPickedNumber * boardInfo.SumOfUnmarkedPositions
}
