package main

import (
	"dive_mod/dive"
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

	commands := strings.Split(string(data), "\n")
	hPosition, depth := dive.ProcessCommands(commands)

	fmt.Println("Part 1")
	fmt.Println("Calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?")
	fmt.Println(hPosition * depth)
	fmt.Println()

	fmt.Println("Part 2")
	hPosition, depth = dive.ProcessCommandsWithAim(commands)
	fmt.Println("Using this new interpretation of the commands, calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?")
	fmt.Println(hPosition * depth)
}
