package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strconv"
)

const BATTERIES int = 12

// func joltageFn(bank string, batteriesLeft int) string {
// 	if len(bank) == 0 {
// 		return ""
// 	}

// 	if len(bank) > batteriesLeft + 1 {
// 		maxIdx := 0
// 		fmt.Println(bank, batteriesLeft, bank[:len(bank)-batteriesLeft])
// 		for i, s := range bank[:len(bank)-batteriesLeft] {
// 			if s > rune(bank[maxIdx]) {
// 				maxIdx = i
// 			}
// 		}
// 	}

// 	return string(bank[maxIdx]) + joltageFn(bank[maxIdx+1:], batteriesLeft-1)
// }

func getLeftMax(str string) int {
	maxIdx := -1
	for i, s := range str {
		if s > rune(maxIdx) {
			maxIdx = i
		}
	}

	return maxIdx
}

func getMaxJoltage(bank string) int {
	// get the max in the digits before BATTERIES
	startIdx := getLeftMax(bank[:len(bank)-BATTERIES+1])
	maxJoltage := string(bank[startIdx])

	result, _ := strconv.Atoi(maxJoltage)

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
