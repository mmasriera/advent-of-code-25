package main

import (
	"fmt"
	"mikicode/aoc25/utils"
)

func isAccessible(rows [][]rune, row int, col int) bool {
	rollCount := 0

	for i := -1; i < 2; i += 1 { // -1 0 1
		for j := -1; j < 2; j += 1 {
			if (row+i >= 0) && (row+i < len(rows)) && (col+j >= 0) && (col+j < len(rows[row])) {
				if rows[row+i][col+j] == '@' {
					rollCount += 1

					if rollCount > 4 { // counting the rows[row][col] too
						return false
					}
				}
			}
		}
	}

	return true
}

func getAccessableRolls(rows [][]rune) int {
	count := 0

	for y, row := range rows {
		for x, item := range row {
			if item == '@' && isAccessible(rows, y, x) {
				count += 1

				rows[y][x] = 'x'
			}
		}
	}

	// if there are accessible rolls, itrerate. again and update matrix
	if count != 0 {
		count += getAccessableRolls(rows)
	}

	return count
}

func main() {
	rows := utils.ReadMatrix(utils.GetInputPath())

	accessableRolls := getAccessableRolls(rows)

	fmt.Println("result:", accessableRolls)
	// test:43 input:9194
}
