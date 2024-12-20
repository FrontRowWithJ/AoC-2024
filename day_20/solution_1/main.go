package main

import (
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := [][]string{}
	startI, startJ := 0, 0
	for i, line := range lines {
		grid = append(grid, strings.Split(line, ""))
		if j := slices.Index(grid[i], "S"); j != -1 {
			startI, startJ = i, j
		}
	}
	path := [][2]int{{startI, startJ}, {startI, startJ}}
	offsets := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for grid[startI][startJ] != "E" {
		for _, offset := range offsets {
			i, j := startI+offset[0], startJ+offset[1]
			if grid[i][j] != "#" && (i != path[len(path)-2][0] || j != path[len(path)-2][1]) {
				path = append(path, [2]int{i, j})
				startI, startJ = i, j
			}
		}
	}
	cheatCount := 0
	for i, behind := range path[1:] {
		for j, ahead := range path[i+2:] {
			manhattan := math.Abs(float64(ahead[1]-behind[1])) + math.Abs(float64(ahead[0]-behind[0]))
			if manhattan <= 20 && (j+1-int(manhattan)) >= 100 {
				cheatCount++
			}
		}
	}
	println(cheatCount)
}
