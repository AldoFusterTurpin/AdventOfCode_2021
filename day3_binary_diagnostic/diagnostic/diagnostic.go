package diagnostic

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessBinaryNumbers(binaryNumbers []string) (gammaRate int64, epsilonRate int64) {
	gammaRateBinary := getMostCommonBit(binaryNumbers)
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

// getZerosCounter returns zerosCounter variable, an slice of ints where zerosCounter[i]
// == the total number of zeros in the index i of
// all the binary numbers of the input slice
// e,g:
// for this input: [00100, 11110, 10110, 10111]
// zerosCounter would be: [1, 3, 0, 1, 3]
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

// getMostCommonBit returns mostCommonBit variable, an string where mostCommonBit[i]:
// a) == 0 if the number of 0s is greater or equal than the number of 1s in the index i of
// all the binary numbers of the input slice
// b) == 0 if the number of 1s is greater than the number of 0s in the index i of
// all the binary numbers of the input slice
func getMostCommonBit(binaryNumbers []string) string {
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
