package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	grid := [][]string{}
	for i, line := range strings.Split(lines[0], "\n") {
		grid = append(grid, make([]string, len(line)))
		for j, cell := range strings.Split(line, "") {
			grid[i][j] = cell
		}
	}
	moves := strings.ReplaceAll(lines[1], "\n", "")
	robotI, robotJ := 0, 0
	for i, line := range grid {
		for j, cell := range line {
			if string(cell) == "@" {
				robotI, robotJ = i, j
				break
			}
		}
	}
	for _, move := range strings.Split(moves, "") {
		if move == "^" {
			for i := robotI; grid[i][robotJ] != "#"; i-- {
				if grid[i][robotJ] == "." {
					for ; i < robotI; i++ {
						grid[i][robotJ] = grid[i+1][robotJ]
					}
					grid[robotI][robotJ] = "."
					robotI--
					break
				}
			}
		} else if move == ">" {
			for j := robotJ; grid[robotI][j] != "#"; j++ {
				if grid[robotI][j] == "." {
					for ; j > robotJ; j-- {
						grid[robotI][j] = grid[robotI][j-1]
					}
					grid[robotI][robotJ] = "."
					robotJ++
					break
				}
			}
		} else if move == "v" {
			for i := robotI; grid[i][robotJ] != "#"; i++ {
				if grid[i][robotJ] == "." {
					for ; i > robotI; i-- {
						grid[i][robotJ] = grid[i-1][robotJ]
					}
					grid[robotI][robotJ] = "."
					robotI++
					break
				}
			}
		} else {
			for j := robotJ; grid[robotI][j] != "#"; j-- {
				if grid[robotI][j] == "." {
					for ; j < robotJ; j++ {
						grid[robotI][j] = grid[robotI][j+1]
					}
					grid[robotI][robotJ] = "."
					robotJ--
					break
				}
			}
		}
	}
	
	total := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "O" {
				total += i*100 + j
			}
		}
	}
	fmt.Println(total)
}
