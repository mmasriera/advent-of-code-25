package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strings"
)

type IdRange struct {
	Min string
	Max string
}

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	var ranges []IdRange

	rangesMode := true
	for _, row := range rows {
		if row == "" {
			rangesMode = false
		} else if rangesMode {
			rangeItems := strings.Split(row, "-")

			ranges = append(ranges, IdRange{rangeItems[0], rangeItems[1]})

			fmt.Println(ranges)
		} else if rangesMode == false {
			fmt.Println(row)
		}
	}

	fmt.Println("result:", rows)
	// test:3 input:
}
