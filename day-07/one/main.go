package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strings"
)

func countSplits(rows []string) int {
	result := 0

	idx := strings.Index(rows[0], "S")

	for i, row := range rows[:1] {
		for j, col := range row {
		}
	}

	fmt.Println(idx)

	return result
}

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	result := countSplits(rows)

	fmt.Println("result:", result)
	// test:21 input:
}
