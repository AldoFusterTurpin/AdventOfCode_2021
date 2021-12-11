package main

import (
	"fmt"
	"io/ioutil"
	"sonar_sweep/measurements"
	"strings"
)

func main() {
	fileName := "input.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file", fileName)
	}

	dataSlice := strings.Split(string(data), "\n")
	dataSliceInts := measurements.ConvertToInts(dataSlice)

	measurementsLargerThanPrevious := measurements.MeasurementsLargerThanPrevious(dataSliceInts)
	windowsMeasurementsLargerThanPrevious := measurements.WindowsMeasurementsLargerThanPrevious(dataSliceInts)

	fmt.Println("Part 1")
	fmt.Println("How many measurements are larger than the previous measurement?")
	fmt.Println(measurementsLargerThanPrevious)

	fmt.Println()

	fmt.Println("Part 2")
	fmt.Println("Consider sums of a three-measurement sliding window. How many sums are larger than the previous sum?")
	fmt.Println(windowsMeasurementsLargerThanPrevious)
}
