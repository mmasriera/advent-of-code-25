package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strconv"
	"strings"
)

// bruteforce -> 0.15s --> TO DO: improve

type IdRange struct {
	Min, Max int
}

func isInRange(rangeList []IdRange, elem int) bool {
	for _, r := range rangeList {
		if (elem >= r.Min) && (elem <= r.Max) {
			return true
		}
	}

	return false
}

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	var ranges []IdRange
	count := 0

	for _, row := range rows {
		if row == "" {
			continue
		}

		elems := strings.Split(row, "-")

		switch len(elems) {
		case 1:
			number, _ := strconv.Atoi(elems[0])
			if isInRange(ranges, number) {
				count += 1
			}

		case 2:
			min, _ := strconv.Atoi(elems[0])
			max, _ := strconv.Atoi(elems[1])
			ranges = append(ranges, IdRange{min, max})
		}
	}

	fmt.Println("result:", count)
	// test:3 input:638
}
