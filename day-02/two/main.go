package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strconv"
	"strings"
)

// string rotation: double the string and find the 2nd index of the1st string before the doubled
// id:9595 -> double:95959595 -> slice 1st:5959595 -> found:5[9595]95 idx=1
func isInvalidId(id string) bool {
	double := id + id
	idx := strings.Index(double[1:], id)

	if idx != -1 && idx < len(id)-1 {
		return true
	}

	return false
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
	// 50857215650
}
