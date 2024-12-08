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
	grid := map[[2]int]bool{}
	var i, j int
	for i = 0; i < len(lines); i++ {
		j = slices.Index(strings.Split(lines[i], ""), "^")
		if j != -1 {
			break
		}
	}
	grid[[2]int{i, j}] = true
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
		} else {
			grid[[2]int{i, j}] = true
		}
	}
	fmt.Println(len(grid))
}
