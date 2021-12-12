package diagnostic

import (
	"fmt"
	"strconv"
	"strings"
)

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

func filterNumbersWithBitToKeepAtIndexI(bitToKeep rune, i int, numbersToFilter []string) []string {
	var filteredElements []string
	for _, number := range numbersToFilter {
		if number[i] == byte(bitToKeep) {
			filteredElements = append(filteredElements, number)
		}
	}
	return filteredElements
}

func getOxygenGeneratorRating(binaryNumbers []string) int64 {
	nBitsPerNumber := len(binaryNumbers[0])

	for i := 0; i < nBitsPerNumber; i++ {
		zerosCounter := getZerosCounter(binaryNumbers)
		nBinaryNumbers := len(binaryNumbers)
		nZerosAtPosI := zerosCounter[i]
		nOnesAtPosI := nBinaryNumbers - nZerosAtPosI

		var bitToKeep rune
		if nZerosAtPosI > nOnesAtPosI {
			bitToKeep = '0'
		} else {
			bitToKeep = '1'
		}

		binaryNumbers = filterNumbersWithBitToKeepAtIndexI(bitToKeep, i, binaryNumbers)
		if len(binaryNumbers) == 1 {
			fmt.Println("filteredElements", binaryNumbers)
			break
		}
	}
	oxygenGeneratorRating, _ := strconv.ParseInt(binaryNumbers[0], 2, 64)
	return oxygenGeneratorRating
}

func getCO2ScrubberRating(binaryNumbers []string) int64 {
	nBitsPerNumber := len(binaryNumbers[0])

	for i := 0; i < nBitsPerNumber; i++ {
		zerosCounter := getZerosCounter(binaryNumbers)
		nBinaryNumbers := len(binaryNumbers)
		nZerosAtPosI := zerosCounter[i]
		nOnesAtPosI := nBinaryNumbers - nZerosAtPosI

		var bitToKeep rune
		if nZerosAtPosI <= nOnesAtPosI {
			bitToKeep = '0'
		} else {
			bitToKeep = '1'
		}

		binaryNumbers = filterNumbersWithBitToKeepAtIndexI(bitToKeep, i, binaryNumbers)
		if len(binaryNumbers) == 1 {
			break
		}
	}
	CO2ScrubberRating, _ := strconv.ParseInt(binaryNumbers[0], 2, 64)
	return CO2ScrubberRating
}

// func getCO2ScrubberRatingOld(binaryNumbers []string) int64 {
// 	nBinaryNumbers := len(binaryNumbers)
// 	zerosCounter := getZerosCounter(binaryNumbers)

// 	for i, nZerosAtPosI := range zerosCounter {
// 		nOnesAtPosI := nBinaryNumbers - nZerosAtPosI

// 		var bitToKeep rune
// 		if nZerosAtPosI <= nOnesAtPosI {
// 			bitToKeep = '0'
// 		} else {
// 			bitToKeep = '1'
// 		}

// 		binaryNumbers = filterNumbersWithBitToKeepAtIndexI(bitToKeep, i, binaryNumbers)
// 		if len(binaryNumbers) == 1 {
// 			break
// 		}
// 	}
// 	CO2ScrubberRating, _ := strconv.ParseInt(binaryNumbers[0], 2, 64)
// 	return CO2ScrubberRating
// }

func GetOxygenGeneratorRatingAndCO2ScrubberRating(binaryNumbers []string) (int64, int64) {
	o := getOxygenGeneratorRating(binaryNumbers)
	co2 := getCO2ScrubberRating(binaryNumbers)
	return o, co2
}
