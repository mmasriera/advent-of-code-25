package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strconv"
)

func isEmptyColumn(matrix [][]rune, col, numRows int) bool {
	for i := 0; i < numRows; i += 1 {
		if matrix[i][col] != ' ' {
			return false
		}
	}

	return true
}

func calculate(numbers []int, operation rune) int {
	result := 0
	if operation == '*' {
		result = 1
	}

	for _, number := range numbers {
		switch operation {
		case '+':
			result += number
		case '*':
			if number != 0 {
				result *= number
			}
		}
	}

	return result
}

func main() {
	matrix := utils.ReadMatrix(utils.GetInputPath())

	rows := len(matrix)
	cols := len(matrix[0])

	var operation rune
	var numbers []int
	result := 0

	for j := range cols {
		if isEmptyColumn(matrix, j, rows) { // end of block (missing last block)
			result += calculate(numbers, operation)
			numbers = nil // reset numbers
		}

		digits := ""

		for i := range rows {
			item := matrix[i][j]

			switch item {
			case ' ':
				continue
			case '*':
				operation = item
			case '+':
				operation = item
			default:
				digits += string(item)
			}
		}

		number, _ := strconv.Atoi(digits)
		numbers = append(numbers, number)
	}

	// last operation
	result += calculate(numbers, operation)

	fmt.Println("result:", result)
	// test:3263827 input:7996218225744
}
