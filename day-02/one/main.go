package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strconv"
	"strings"
)

func isInvalidId(id string) bool {
	half := len(id) / 2

	return id[:half] == id[half:]
}

func getInvalidIdsInRange(start int, end int) int {
	var result = 0

	for start <= end {
		if isInvalidId(strconv.Itoa(start)) {
			result += start
		}

		start += 1
	}

	return result
}

func main() {
	lines := utils.ReadLines(utils.GetInputPath())
	idRanges := strings.Split(lines[0], ",")
	result := 0

	for _, idRange := range idRanges {
		ids := strings.Split(idRange, "-")

		start, _ := strconv.Atoi(ids[0])
		end, _ := strconv.Atoi(ids[1])

		invalidsCount := getInvalidIdsInRange(start, end)

		result += invalidsCount
	}

	fmt.Println("result:", result)
	// 40055209690
}
