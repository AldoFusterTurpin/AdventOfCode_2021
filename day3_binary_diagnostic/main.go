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

	gammaRate, epsilonRate := diagnostic.ProcessBinaryNumbers(binaryNumbers)
	fmt.Println("Part 1:")
	fmt.Println("Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate, then multiply them together. What is the power consumption of the submarine? (Be sure to represent your answer in decimal, not binary.")
	fmt.Println("gammaRate is:", gammaRate)
	fmt.Println("epsilonRate is:", epsilonRate)
	fmt.Println("gammaRate * epsilonRate is:", gammaRate*epsilonRate)
}
