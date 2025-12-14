package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadMatrix(path string) [][]rune {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]rune

	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	return matrix
}

const usageExample string = "usage: go run main.go test.txt"

func GetInputPath() string {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("missing test file:\n%s", usageExample))
	}

	input := os.Args[1]

	if !strings.HasSuffix(input, ".txt") {
		input += ".txt"
	}

	return fmt.Sprintf("../inputs/%s", input)
}
