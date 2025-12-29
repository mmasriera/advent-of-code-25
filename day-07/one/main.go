package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strings"
)

type Position struct {
	Row, Col int
}

func appendUniquePosition(list []Position, item Position) []Position {
	for _, pos := range list {
		if (pos.Row == item.Row) && (pos.Col == item.Col) {
			return list
		}
	}

	return append(list, item)
}

func countSplits(rows []string) int {
	count := 0
	var items []Position // TO DO: use only the column idx --> []Position vs []int

	for i := range len(rows) {
		if i == 0 {
			startIdx := strings.Index(rows[0], "S")
			items = append(items, Position{Row: 1, Col: startIdx})

			continue
		}

		var newItems []Position
		for _, item := range items {
			if string(rows[item.Row][item.Col]) == "^" { // split -> check if split in border?
				count += 1
				newItems = appendUniquePosition(newItems, Position{Row: item.Row + 1, Col: item.Col - 1})
				newItems = appendUniquePosition(newItems, Position{Row: item.Row + 1, Col: item.Col + 1})

			} else {
				newItems = appendUniquePosition(newItems, Position{Row: item.Row + 1, Col: item.Col})
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
