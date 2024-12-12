package main

import (
	"os"
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
		perimeter := 0
		for _, point := range points {
			for _, direction := range directions {
				i, j := point.i+direction.i, point.j+direction.j
				if i < 0 || i >= len(lines) || j < 0 || j >= len(lines) {
					perimeter++
				} else if lines[i][j] != lines[point.i][point.j] {
					perimeter++
				}
			}
		}
		price += len(points) * perimeter
	}
	println(price)
}
