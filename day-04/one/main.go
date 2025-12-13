package main

import (
	"fmt"
	"mikicode/aoc25/utils"
)

func getAccessableRolls(rows []string) int {
	count := 0

	for y, row := range rows {
		for x, item := range row {
			fmt.Println(y, x, item == '.')
		}
	}

	return count
}

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	accessableRolls := getAccessableRolls(rows)

	fmt.Println("result:", accessableRolls)
}
