package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	antennas := map[string][]Point{}
	grid := map[Point]bool{}
	for i, line := range lines {
		for j, cell := range strings.Split(line, "") {
			if cell != "." {
				antennas[cell] = append(antennas[cell], Point{i, j})
			}
		}
	}
	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				a, b := points[i], points[j]
				x, y := a.X-b.X, a.Y-b.Y
				grid[points[i]], grid[points[j]] = true, true
				for x1, y1 := a.X+x, a.Y+y; x1 >= 0 && y1 >= 0 && x1 < len(lines) && y1 < len(lines); x1, y1 = x1+x, y1+y {
					grid[Point{x1, y1}] = true
				}
				for x2, y2 := b.X-x, b.Y-y; x2 >= 0 && y2 >= 0 && x2 < len(lines) && y2 < len(lines); x2, y2 = x2-x, y2-y {
					grid[Point{x2, y2}] = true
				}
			}
		}
	}
	fmt.Println(len(grid))
}
