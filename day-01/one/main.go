package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strconv"
)

func main() {
	var zeroSum int = 0
	var position int = 50
	lines := utils.ReadLines(utils.GetInputPath())

	for _, line := range lines {
		direction := line[0:1]
		value, _ := strconv.Atoi(line[1:])
		value %= 100

		if direction == "L" {
			position -= value

			if position < 0 {
				position = 100 + position
			}
		} else if direction == "R" {
			position += value

			if position > 99 {
				position = position - 100
			}
		}

		if position == 0 {
			zeroSum++
		}
	}

	fmt.Println("result:", zeroSum)
	// test.txt -> 3
	// input.txt -> 1084
}
