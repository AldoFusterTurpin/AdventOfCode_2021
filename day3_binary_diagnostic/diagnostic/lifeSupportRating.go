package diagnostic

import "strconv"

type BitKeeper interface {
	getBitToKeep(int, int) rune
}

type Oxygen struct {
}

func (o Oxygen) getBitToKeep(nZerosAtPosI, nOnesAtPosI int) rune {
	if nZerosAtPosI > nOnesAtPosI {
		return '0'
	}
	return '1'
}

type CO2Scrubber struct {
}

func (co2 CO2Scrubber) getBitToKeep(nZerosAtPosI, nOnesAtPosI int) rune {
	if nZerosAtPosI <= nOnesAtPosI {
		return '0'
	}
	return '1'
}

func GetOxygenGeneratorRatingAndCO2ScrubberRating(binaryNumbers []string) (int64, int64) {
	oxygenBitKeeper := Oxygen{}
	CO2ScrubberBitKeeper := CO2Scrubber{}

	oxygenGeneratorRating := getResult(binaryNumbers, oxygenBitKeeper)
	cO2ScrubberRating := getResult(binaryNumbers, CO2ScrubberBitKeeper)
	return oxygenGeneratorRating, cO2ScrubberRating
}

func getResult(binaryNumbers []string, bk BitKeeper) int64 {
	nBitsPerNumber := len(binaryNumbers[0])

	for i := 0; i < nBitsPerNumber; i++ {
		zerosCounter := getZerosCounter(binaryNumbers)
		nZerosAtPosI := zerosCounter[i]
		nBinaryNumbers := len(binaryNumbers)
		nOnesAtPosI := nBinaryNumbers - nZerosAtPosI

		bitToKeep := bk.getBitToKeep(nZerosAtPosI, nOnesAtPosI)

		binaryNumbers = filterNumbersWithBitToKeepAtIndexI(bitToKeep, i, binaryNumbers)
		if len(binaryNumbers) == 1 {
			break
		}
	}
	result, _ := strconv.ParseInt(binaryNumbers[0], 2, 64)
	return result
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
