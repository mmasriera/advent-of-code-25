package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strings"
)

func appendUniquePosition(list []int, item int) []int {
	for _, pos := range list {
		if item == pos {
			return list
		}
	}

	return append(list, item)
}

func countSplits(rows []string) int {
	count := 0
	var items []int

	for i := range len(rows) {
		if i == 0 {
			startIdx := strings.Index(rows[0], "S")
			items = append(items, startIdx)

			continue
		}

		var newItems []int
		for _, item := range items {
			if string(rows[i][item]) == "^" { // split -> check if split in border?
				count += 1
				newItems = appendUniquePosition(newItems, item-1)
				newItems = appendUniquePosition(newItems, item+1)

			} else {
				newItems = appendUniquePosition(newItems, item)
			}
		}

		items = newItems
	}

	return count
}

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	result := countSplits(rows)

	fmt.Println("result:", result)
	// test:21 input:1587
}
