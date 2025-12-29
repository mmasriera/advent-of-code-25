package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strings"
)

func countSplits(rows []string) int {
	var items []int

	for i := range len(rows) {
		fmt.Println(i, len(items))
		if i == 0 {
			startIdx := strings.Index(rows[0], "S")
			items = append(items, startIdx)

			continue
		}

		var newItems []int
		for _, item := range items {
			if string(rows[i][item]) == "^" { // split -> check if split in border?
				newItems = append(newItems, item-1)
				newItems = append(newItems, item+1)

			} else {
				newItems = append(newItems, item)
			}
		}

		items = newItems
	}

	return len(items)
}

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	result := countSplits(rows)

	fmt.Println("result:", result)
	// test:21 input:1587
}
