package main

import (
	"cmp"
	"fmt"
	"math"
	"mikicode/aoc25/utils"
	"slices"
	"strconv"
	"strings"
)

type Point3d struct {
	X, Y, Z int
}

// type Pairs map[string]map[string]float64
type Pair struct {
	P1, P2 string
	Dist   float64
}

func make3dPoint(point string) Point3d {
	splitted := strings.Split(point, ",")
	x, _ := strconv.Atoi(splitted[0])
	y, _ := strconv.Atoi(splitted[1])
	z, _ := strconv.Atoi(splitted[2])

	return Point3d{X: x, Y: y, Z: z}
}

func distance(point1, point2 string) float64 {
	p1 := make3dPoint(point1)
	p2 := make3dPoint(point2)

	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2) + math.Pow(float64(p2.Z-p1.Z), 2))
}

func getCombinationsCount(n int) int {
	return (n * (n / 2)) - (n / 2)
}

func buildPairsWithDistance(points []string) []Pair {
	pairs := make([]Pair, 0, getCombinationsCount(len(points)))

	for _, point1 := range points {
		for _, point2 := range points {
			if point1 == point2 {
				continue
			}

			// it is way faster to (add all + sort + compact) than check here if the pair is in the list
			pairs = append(pairs, Pair{P1: point1, P2: point2, Dist: distance(point1, point2)})
		}
	}

	slices.SortFunc(pairs, func(a, b Pair) int {
		return cmp.Compare(a.Dist, b.Dist)
	})

	// remove consecutive repetitions
	return slices.CompactFunc(pairs, func(a, b Pair) bool {
		return a.Dist == b.Dist
	})
}

func makeCircuits(pairs []Pair) [][]Pair {
	var circuits [][]Pair

	for _, pair := range pairs {
		idx := slices.IndexFunc(circuits, func(c []Pair) bool {
			for _, p := range c {
				return (p.P1 == pair.P1) || (p.P2 == pair.P2) || (p.P1 == pair.P2) || (p.P2 == pair.P1)
			}

			return false
		})

		if idx != -1 {
			circuits[idx] = append(circuits[idx], pair)
		} else {
			circuits = append(circuits, []Pair{pair})
		}

	}

	return circuits
}

func countNlargest(circuits [][]Pair, n int) int {
	var sizes []int

	for _, c := range circuits {
		sizes = append(sizes, len(c))
	}

	slices.SortFunc(sizes, func(a, b int) int {
		return b - a // sort from higher to lower
	})

	fmt.Println("sizes:", sizes)

	result := 1

	for _, pairCount := range sizes[:n] {
		result *= pairCount + 1
	}

	return result
}

func getLargestCircuits(points []string, connections int) int {
	pairs := buildPairsWithDistance(points)

	circuits := makeCircuits(pairs[:connections])

	result := countNlargest(circuits, 3)

	return result
}

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	largestCircuits := getLargestCircuits(rows, 1000)

	fmt.Println("result:", largestCircuits)
	// test:40 input:990 too low
}
