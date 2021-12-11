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

func flipBit(c byte) byte {
	if c == '0' {
		return '1'
	}
	return '0'
}

func flipBits(binNum string) []byte {
	flippedBits := make([]byte, len(binNum))

	for i, v := range binNum {
		flippedBits[i] = flipBit(byte(v))
	}
	return flippedBits
}

func getZerosCounter(binaryNumbers []string) []int {
	nBitsPerNumber := len(binaryNumbers[0])
	zerosCounter := make([]int, nBitsPerNumber)

	for _, binaryNumber := range binaryNumbers {
		for i, bit := range binaryNumber {
			if bit == '0' {
				zerosCounter[i]++
			}
		}
	}
	return zerosCounter
}

func getGammaRateBinary(binaryNumbers []string) string {
	zerosCounter := getZerosCounter(binaryNumbers)
	nBinaryNumbers := len(binaryNumbers)
	var mostCommonBit strings.Builder

	for _, nZerosAtPosI := range zerosCounter {
		nOnesAtPosI := nBinaryNumbers - nZerosAtPosI

		if nZerosAtPosI >= nOnesAtPosI {
			mostCommonBit.WriteString("0")
		} else {
			mostCommonBit.WriteString("1")
		}
	}

	return mostCommonBit.String()
}

func ProcessBinaryNumbers(binaryNumbers []string) (gammaRate int64, epsilonRate int64) {
	gammaRateBinary := getGammaRateBinary(binaryNumbers)
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
