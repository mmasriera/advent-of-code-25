package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strings"
)

type PositionCount struct {
	Col, Value int
}

func countStep(step []PositionCount) int {
	result := 0

	for _, item := range step {
		result += item.Value
	}

	return result
}

func appendWithSum(items []PositionCount, col int, carry int) []PositionCount {
	for i, item := range items {
		if item.Col == col {
			items[i].Value += carry

			return items
		}
	}

	return append(items, PositionCount{Col: col, Value: carry})
}

// in each step, summ the overlapping paths
func countPaths(rows []string) int {
	var step []PositionCount

	for i := range len(rows) {
		if i == 0 {
			startIdx := strings.Index(rows[0], "S")
			step = append(step, PositionCount{Col: startIdx, Value: 1})

			continue
		}

		var newStep []PositionCount
		for _, item := range step {
			if string(rows[i][item.Col]) == "^" {
				newStep = appendWithSum(newStep, item.Col-1, item.Value)
				newStep = appendWithSum(newStep, item.Col+1, item.Value)
			} else {
				newStep = appendWithSum(newStep, item.Col, item.Value)
			}
		}

		step = newStep
	}

	return countStep(step)
}

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	result := countPaths(rows)

	fmt.Println("result:", result)
	// test:40 input:5748679033029
}
