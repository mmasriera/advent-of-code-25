package main

// https://algs4.cs.princeton.edu/15uf/
// https://www.geeksforgeeks.org/dsa/introduction-to-disjoint-set-data-structure-or-union-find-algorithm/

// https://www.w3schools.com/dsa/dsa_algo_mst_kruskal.php

import (
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

			exists := slices.ContainsFunc(pairs, func(p Pair) bool {
				return (p.P1 == point2) && (p.P2 == point1)
			})

			if !exists {
				pairs = append(pairs, Pair{P1: point1, P2: point2, Dist: distance(point1, point2)})
			}
		}
	}

	return pairs
}

func getLargestCircuits(points []string, connections int) []int {
	pairs := buildPairsWithDistance(points)

	fmt.Println(pairs)

	return []int{len(pairs)}
}

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	largestCircuits := getLargestCircuits(rows, 10)

	fmt.Println("result:", largestCircuits)
	// test:40 input:
}
