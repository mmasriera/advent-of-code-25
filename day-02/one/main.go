package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strings"
)

func main() {
	lines := utils.ReadLines(utils.GetInputPath())
	idRanges := strings.Split(lines[0], ",")

	for _, idRange := range idRanges {
		ids := strings.Split(idRange, "-")

		fmt.Println(ids)
	}
}
