package main

import (
	"binary_diagnostic/diagnostic"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileName := "input.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file", fileName)
	}

	binaryNumbers := strings.Split(string(data), "\n")

	fmt.Println("Part 1:")
	gammaRate, epsilonRate := diagnostic.ProcessBinaryNumbers(binaryNumbers)
	fmt.Println("Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate, then multiply them together. What is the power consumption of the submarine?")
	fmt.Println("gammaRate is:", gammaRate)
	fmt.Println("epsilonRate is:", epsilonRate)
	fmt.Println("gammaRate * epsilonRate is:", gammaRate*epsilonRate)

	fmt.Println("Part 2:")
	oxygenGeneratorRating, CO2ScrubberRating := diagnostic.GetOxygenGeneratorRatingAndCO2ScrubberRating(binaryNumbers)
	fmt.Println("Use the binary numbers in your diagnostic report to calculate the oxygen generator rating and CO2 scrubber rating, then multiply them together. What is the life support rating of the submarine?")
	fmt.Println("oxygenGeneratorRating is:", oxygenGeneratorRating)
	fmt.Println("CO2ScrubberRating is:", CO2ScrubberRating)
	fmt.Println("oxygenGeneratorRating * epsilonRate is:", oxygenGeneratorRating*CO2ScrubberRating)
}
