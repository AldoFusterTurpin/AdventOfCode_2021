package dive

import (
	"strconv"
	"strings"
)

const (
	Forward = "forward"
	Down    = "down"
	Up      = "up"
)

func ProcessCommands(commands []string) (int, int) {
	horizontalPosition := 0
	depth := 0
	for _, c := range commands {
		s := strings.Split(c, " ")
		command, unitsStr := s[0], s[1]
		units, _ := strconv.Atoi(unitsStr)

		switch command {
		case Forward:
			horizontalPosition += units
		case Down:
			depth += units
		case Up:
			depth -= units
		}
	}
	return horizontalPosition, depth
}

// Part2
func ProcessCommandsWithAim(commands []string) (int, int) {
	horizontalPosition := 0
	depth := 0
	aim := 0
	for _, c := range commands {
		s := strings.Split(c, " ")
		command, unitsStr := s[0], s[1]
		units, _ := strconv.Atoi(unitsStr)

		switch command {
		case Forward:
			horizontalPosition += units
			depth += aim * units
		case Down:
			aim += units
		case Up:
			aim -= units
		}
	}
	return horizontalPosition, depth
}
