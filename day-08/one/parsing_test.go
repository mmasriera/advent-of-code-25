package main

import (
	"fmt"
	"mikicode/aoc25/utils"
	"strconv"
	"strings"
	"testing"
)

var rows = utils.ReadLines("../inputs/test.txt")

func BenchmarkManual(b *testing.B) {
	for b.Loop() {
		points := make([]Point3d, 0, len(rows))
		for _, row := range rows {
			splitted := strings.Split(row, ",")
			x, _ := strconv.Atoi(splitted[0])
			y, _ := strconv.Atoi(splitted[1])
			z, _ := strconv.Atoi(splitted[2])
			points = append(points, Point3d{X: x, Y: y, Z: z})
		}
	}
}

func BenchmarkSscanf(b *testing.B) {
	for b.Loop() {
		points := make([]Point3d, 0, len(rows))
		for _, row := range rows {
			var p Point3d
			fmt.Sscanf(row, "%d,%d,%d", &p.X, &p.Y, &p.Z)
			points = append(points, p)
		}
	}
}
