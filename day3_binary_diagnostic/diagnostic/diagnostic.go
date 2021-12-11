package diagnostic

import (
	"fmt"
	"strconv"
	"strings"
)

/*
zerosCounter[i] returns the total number of zeros in the index i of
all the binary numbers of the input data
e,g:
for this input:
00100
11110
10110
10111
zerosCounter would be:
[1, 3, 0, 1, 3]
*/

// func convertSliceOfIntsToStr(input []int) string {
// 	sliceOfStr := make([]string, len(input))
// 	for i, v := range input {
// 		sliceOfStr[i] = strconv.Itoa(v)
// 	}
// 	return strings.Join(sliceOfStr, "")
// }

// mostCommonBit[i] returns the most  common bit in the index i of all the binary
// numbers of the input data

func flipBit(c rune) byte {
	if c == '0' {
		return '1'
	}
	return '0'
}

func flipBits(input string) []byte {
	flippedBits := make([]byte, len(input))

	for i, v := range input {
		flippedBits[i] = flipBit(v)
	}
	return flippedBits
}

func ProcessBinaryNumbers(binaryNumbers []string) (gammaRate int64, epsilonRate int64) {
	nBitsPerNumber := len(binaryNumbers[0])
	zerosCounter := make([]int, nBitsPerNumber)
	var mostCommonBit strings.Builder

	for _, binaryNumber := range binaryNumbers {
		for i, bit := range binaryNumber {
			if bit == '0' {
				zerosCounter[i]++
			}
		}
	}

	nBinaryNumbers := len(binaryNumbers)

	for _, nZerosAtPosI := range zerosCounter {
		nOnesAtPosI := nBinaryNumbers - nZerosAtPosI

		if nZerosAtPosI >= nOnesAtPosI {
			mostCommonBit.WriteString("0")
		} else {
			mostCommonBit.WriteString("1")
		}
	}

	gammaRateBinary := mostCommonBit.String()

	gammaRate, err := strconv.ParseInt(gammaRateBinary, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	epsilonRateBinary := flipBits(gammaRateBinary)

	epsilonRate, err = strconv.ParseInt(string(epsilonRateBinary), 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	return gammaRate, epsilonRate
}
