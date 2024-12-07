package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := make([][]bool, len(lines))
	for i := range grid {
		grid[i] = make([]bool, len(lines[0]))
	}
	var i, j int
	for i = 0; i < len(lines); i++ {
		j = slices.Index(strings.Split(lines[i], ""), "^")
		if j != -1 {
			break
		}
	}
	totalUnique := 1
	grid[i][j] = true
	directionI := []int{-1, 0, 1, 0}
	directionJ := []int{0, 1, 0, -1}
	directionIndex := 0
	for {
		i, j = i+directionI[directionIndex], j+directionJ[directionIndex]
		if i == -1 || i == len(lines) || j == -1 || j == len(lines[0]) {
			break
		}
		if string(lines[i][j]) == "#" {
			i, j = i-directionI[directionIndex], j-directionJ[directionIndex]
			directionIndex = (directionIndex + 1) % 4
		} else if !grid[i][j] {
			grid[i][j] = true
			totalUnique++
		}
	}
	fmt.Println(totalUnique)
}
