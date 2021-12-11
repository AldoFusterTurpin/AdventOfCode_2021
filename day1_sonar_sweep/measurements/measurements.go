package measurements

import (
	"strconv"
)

func ConvertToInts(data []string) []int {
	result := make([]int, len(data))
	for i, x := range data {
		converted, _ := strconv.Atoi(x)
		result[i] = converted
	}
	return result
}

func MeasurementsLargerThanPrevious(data []int) int {
	i := 1
	dataLenght := len(data)
	measurementsLargerThanPrevious := 0

	for i < dataLenght {
		if data[i-1] < data[i] {
			measurementsLargerThanPrevious++
		}
		i++
	}
	return measurementsLargerThanPrevious
}

// Part 2
func CreateSlidingWindowSlice(data []int) []int {
	slidingWindowSlice := make([]int, 0)
	i := 2
	for i < len(data) {
		slidingWindowSlice = append(slidingWindowSlice, data[i]+data[i-1]+data[i-2])
		i++
	}
	return slidingWindowSlice
}

func WindowsMeasurementsLargerThanPrevious(data []int) int {
	slidingWindowSlice := CreateSlidingWindowSlice(data)
	return MeasurementsLargerThanPrevious(slidingWindowSlice)
}
