package main

// https://algs4.cs.princeton.edu/15uf/
// https://www.geeksforgeeks.org/dsa/introduction-to-disjoint-set-data-structure-or-union-find-algorithm/

// https://www.w3schools.com/dsa/dsa_algo_mst_kruskal.php

import (
	"fmt"
	"math"
	"mikicode/aoc25/utils"
	"strconv"
	"strings"
)

type Point3d struct {
	X, Y, Z int
}

type Pairs map[string]map[string]float64

func strToPoint(str string) Point3d {
	splitted := strings.Split(str, ",")
	x, _ := strconv.Atoi(splitted[0])
	y, _ := strconv.Atoi(splitted[1])
	z, _ := strconv.Atoi(splitted[2])

	return Point3d{X: x, Y: y, Z: z}
}

func distance(point1, point2 string) float64 {
	p1 := strToPoint(point1)
	p2 := strToPoint(point2)

	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2) + math.Pow(float64(p2.Z-p1.Z), 2))
}

func buildPairsMap(points []string) Pairs {
	result := make(Pairs)

	for _, point := range points {
		for j := range points {
			point2 := points[j]
			if _, ok := result[point2]; ok == false && point != point2 {
				pair := map[string]float64{
					point2: distance(point, point2),
				}
				result[point] = pair
			}
		}
	}

	return result
}

func main() {
	rows := utils.ReadLines(utils.GetInputPath())
	// map[point1][point2:distance]
	pairsMap := buildPairsMap(rows)

	fmt.Println("result:", pairsMap)
	// test:40 input:
}
