package main

import (
	"os"
	"slices"
	"strings"
)

type Point struct {
	i int
	j int
}

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := make([][]bool, len(lines))
	for i, line := range lines {
		grid[i] = make([]bool, len(line))
	}
	regions := [][]Point{}
	directions := [4]Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i, line := range lines {
		for j, cell := range line {
			if grid[i][j] {
				continue
			}
			regions = append(regions, []Point{{i, j}})
			grid[i][j] = true
			stack := []Point{{i, j}}
			for len(stack) != 0 {
				point := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				for _, direction := range directions {
					x, y := point.i+direction.i, point.j+direction.j
					if x >= 0 && y >= 0 && x < len(lines) && y < len(lines) {
						if !grid[x][y] && rune(lines[x][y]) == cell {
							grid[x][y] = true
							regions[len(regions)-1] = append(regions[len(regions)-1], Point{x, y})
							stack = append(stack, Point{x, y})
						}
					}
				}
			}
		}
	}
	price := 0
	for _, points := range regions {
		tops, lefts, rights, bottoms := []Point{}, []Point{}, []Point{}, []Point{}
		for _, point := range points {
			if point.i-1 == -1 || lines[point.i-1][point.j] != lines[point.i][point.j] {
				tops = append(tops, point)
			}
			if point.j-1 == -1 || lines[point.i][point.j-1] != lines[point.i][point.j] {
				lefts = append(lefts, point)
			}
			if point.i+1 == len(lines) || lines[point.i+1][point.j] != lines[point.i][point.j] {
				bottoms = append(bottoms, point)
			}
			if point.j+1 == len(lines) || lines[point.i][point.j+1] != lines[point.i][point.j] {
				rights = append(rights, point)
			}
		}
		slices.SortFunc(tops, func(a Point, b Point) int { return (a.i*len(lines) + a.j) - (b.i*len(lines) + b.j) })
		slices.SortFunc(bottoms, func(a Point, b Point) int { return (a.i*len(lines) + a.j) - (b.i*len(lines) + b.j) })
		slices.SortFunc(rights, func(a Point, b Point) int { return (a.j*len(lines) + a.i) - (b.j*len(lines) + b.i) })
		slices.SortFunc(lefts, func(a Point, b Point) int { return (a.j*len(lines) + a.i) - (b.j*len(lines) + b.i) })
		numOfSides := 4
		currI := tops[0].i
		for idx := 0; idx < len(tops)-1; idx++ {
			if tops[idx+1].j != tops[idx].j+1 || currI != tops[idx+1].i {
				numOfSides++
				currI = tops[idx+1].i
			}
		}

		currI = bottoms[0].i
		for idx := 0; idx < len(bottoms)-1; idx++ {
			if bottoms[idx+1].j != bottoms[idx].j+1 || currI != bottoms[idx+1].i {
				numOfSides++
				currI = bottoms[idx+1].i
			}
		}
		currJ := lefts[0].j
		for idx := 0; idx < len(lefts)-1; idx++ {
			if lefts[idx+1].i != lefts[idx].i+1 || currJ != lefts[idx+1].j {
				numOfSides++
				currJ = lefts[idx+1].j
			}
		}
		currJ = rights[0].j
		for idx := 0; idx < len(rights)-1; idx++ {
			if rights[idx+1].i != rights[idx].i+1 || currJ != rights[idx+1].j {
				numOfSides++
				currJ = rights[idx+1].j
			}
		}
		price += len(points) * numOfSides
	}
	println(price)
}
