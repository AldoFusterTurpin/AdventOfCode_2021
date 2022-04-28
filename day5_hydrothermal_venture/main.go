package main

import (
	"bufio"
	"fmt"
	"hydrothermalventure/internal/detector"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := make([]detector.PairCoordinates, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pc := buildPairCoordinates(line)
		s = append(s, pc)
	}

	rb := detector.HorizontalAndVerticalLinesBuilder{}
	fmt.Println(detector.GetNumberOfOverlappingPoints(s, 2, rb))
}

func buildPairCoordinates(line string) detector.PairCoordinates {
	parts := strings.Split(line, "->")
	l := strings.Trim(parts[0], "")
	r := strings.Trim(parts[1], "")

	pc := detector.PairCoordinates{
		InitCoordinate: buildCoordinate(l),
		EndCoordinate:  buildCoordinate(r),
	}
	return pc
}

// buildCoordinate, given a string like "957,596", return the corresponding
// Coordinate variable
func buildCoordinate(s string) detector.Coordinate {
	splitted := strings.Split(s, ",")
	x := strings.Trim(splitted[0], " ")
	y := strings.Trim(splitted[1], " ")

	X, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}

	Y, err := strconv.Atoi(y)
	if err != nil {
		panic(err)
	}
	return detector.Coordinate{
		X: X,
		Y: Y,
	}
}
