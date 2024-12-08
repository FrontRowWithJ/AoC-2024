package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	antennas := map[string][][2]int{}
	grid := map[[2]int]bool{}
	for i, line := range lines {
		for j, cell := range strings.Split(line, "") {
			if cell != "." {
				antennas[cell] = append(antennas[cell], [2]int{i, j})
			}
		}
	}
	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				a, b := points[i], points[j]
				x, y := a[0]-b[0], a[1]-b[1]
				grid[points[i]], grid[points[j]] = true, true
				for x1, y1 := a[0]+x, a[1]+y; x1 >= 0 && y1 >= 0 && x1 < len(lines) && y1 < len(lines); x1, y1 = x1+x, y1+y {
					grid[[2]int{x1, y1}] = true
				}
				for x2, y2 := b[0]-x, b[1]-y; x2 >= 0 && y2 >= 0 && x2 < len(lines) && y2 < len(lines); x2, y2 = x2-x, y2-y {
					grid[[2]int{x2, y2}] = true
				}
			}
		}
	}
	fmt.Println(len(grid))
}
