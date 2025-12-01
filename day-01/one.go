
package main

import (
	"mikicode/aoc25/utils"
	"strconv"
	"fmt"
)

func main() {
	var zeroSum int = 0
	var position int = 50
	lines := utils.ReadLines("./inputs/one.txt")

	for _, line := range lines {
		direction := line[0:1]
		rawValue, _ := strconv.Atoi(line[1:])
		value := rawValue % 100

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
	// test -> 3
	// one -> 1084
}
