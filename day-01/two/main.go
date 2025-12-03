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

		if value > 99 {
			zeroSum += value / 100
			value %= 100
		}

		var nextPosition int

		if direction == "L" {
			nextPosition = position - value
		} else if direction == "R" {
			nextPosition = position + value
		}

		if nextPosition < 0 {
			nextPosition = 100 + nextPosition
			if position != 0 {
				zeroSum++
			}
		} else if nextPosition > 99 {
			nextPosition = nextPosition - 100
			if position != 0 {
				zeroSum++
			}
		} else if position != 0 && nextPosition == 0 {
			zeroSum++
		}

		position = nextPosition
	}

	fmt.Println("result:", zeroSum)
	// test.txt -> 6
	// input.txt -> 6475
}
