package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strconv"
)

const BATTERIES int = 12

func getLeftMax(str string) int {
	maxIdx := 0
	for i, s := range str {
		if s > rune(str[maxIdx]) {
			maxIdx = i
		}
	}

	return maxIdx
}

/*
get the max value between bank[0] and bank[batteries-1] -> 1st digit at index "maxIdx"
then from the next bank[maxIdx+1] to  bank[batteries-2], do the same -> 2nd digit and update "maxIds"
iterate
*/
func getMaxJoltage(bank string, batteriesLeft int) string {
	if batteriesLeft == 0 {
		return ""
	} else if len(bank) == batteriesLeft {
		return bank
	}

	maxIdx := getLeftMax(bank[:len(bank)-batteriesLeft+1])

	return string(bank[maxIdx]) + getMaxJoltage(bank[maxIdx+1:], batteriesLeft-1)
}

func main() {
	banks := utils.ReadLines(utils.GetInputPath())
	result := 0

	for _, bank := range banks {
		maxJolatage := getMaxJoltage(bank, BATTERIES)
		maxJolatageValue, _ := strconv.Atoi(maxJolatage)

		result += maxJolatageValue
	}

	fmt.Println("result:", result)
	// test:3121910778619 input:175304218462560
}
