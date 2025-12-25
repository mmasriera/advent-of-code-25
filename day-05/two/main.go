package main

import (
	"cmp"
	"fmt"
	"mikicode/aoc25/utils"
	"slices"
	"strconv"
	"strings"
)

type IdRange struct {
	Min, Max int
}

func countFreshIds(ranges []IdRange) int {
	count := 0

	slices.SortFunc(ranges, func(a, b IdRange) int {
		return cmp.Compare(a.Min, b.Min)
	})

	prevMax := 0

	for _, r := range ranges {
		if prevMax > r.Max {
			continue
		}

		if prevMax >= r.Min { // overlapping -> remove the overlap
			count += r.Max - prevMax
		} else {
			count += r.Max - r.Min + 1 // no overlapping -> count the starting item
		}

		prevMax = r.Max
	}

	return count
}

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	var ranges []IdRange

	for _, row := range rows {
		if row == "" {
			continue
		}

		elems := strings.Split(row, "-")

		if len(elems) == 2 { // only ranges
			min, _ := strconv.Atoi(elems[0])
			max, _ := strconv.Atoi(elems[1])
			ranges = append(ranges, IdRange{min, max})
		}
	}

	count := countFreshIds(ranges)

	fmt.Println("result:", count)
	// test:14 input:352946349407338
}
