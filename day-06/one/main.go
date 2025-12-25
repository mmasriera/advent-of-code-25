package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strconv"
	"strings"
)

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	var grid [][]string

	for _, row := range rows {
		grid = append(grid, strings.Fields(row))
	}

	numItems := len(grid)
	operations := grid[numItems-1] // last row
	result := 0

	for col, operation := range operations {
		acc := 0

		if operation == "*" {
			acc = 1
		}

		for i := range numItems - 1 {
			number, _ := strconv.Atoi(grid[i][col])

			switch operation {
			case "+":
				acc += number
			case "*":
				acc *= number
			}
		}

		result += acc
	}

	fmt.Println("result:", result)
	// test:4277556 input:5784380717354
}
