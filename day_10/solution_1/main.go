package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := make([][]int, len(lines))
	stack := [][2]int{}
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, cell := range strings.Split(line, "") {
			grid[i][j] = map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}[cell]
			if grid[i][j] == 0 {
				stack = append(stack, [2]int{i, j})
			}
		}
	}
	count := 0
	directions := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for len(stack) != 0 {
		point := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		pos := grid[point[0]][point[1]]
		if pos == 9 {
			count++
		}
		for _, direction := range directions {
			i, j := point[0]+direction[0], point[1]+direction[1]
			if i >= 0 && j >= 0 && i < len(grid) && j < len(grid[i]) && grid[i][j] == pos+1 {
				stack = append(stack, [2]int{i, j})
			}
		}
	}
	fmt.Println(count)
}
