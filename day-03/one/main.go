package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strconv"
)

func getMaxJoltage(bank string) int {
	idx1 := 0

	// get highest value (except last position)
	for i, s := range bank[:len(bank)-1] {
		if s > rune(bank[idx1]) {
			idx1 = i
		}
	}

	// get highest value (starting from the last idx's next index)
	offset := idx1 + 1
	idx2 := offset
	for i, s := range bank[idx2:] {
		if s > rune(bank[idx2]) {
			idx2 = offset + i
		}
	}

	joltageStr := bank[idx1:idx1+1] + bank[idx2:idx2+1] // bytes -> strings
	result, _ := strconv.Atoi(joltageStr)

	return result
}

func main() {
	banks := utils.ReadLines(utils.GetInputPath())
	result := 0

	for _, bank := range banks {
		maxJolatage := getMaxJoltage(bank)

		result += maxJolatage
	}

	fmt.Println("result:", result)
	// test:357 input:17613
}
