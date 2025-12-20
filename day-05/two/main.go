package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strconv"
	"strings"
)

type IdRange struct {
	Min, Max int
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

		if len(elems) == 2 {
			min, _ := strconv.Atoi(elems[0])
			max, _ := strconv.Atoi(elems[1])
			ranges = append(ranges, IdRange{min, max})
		}
	}

	fmt.Println("result:", count, ranges)
	// test:14 input:
}
