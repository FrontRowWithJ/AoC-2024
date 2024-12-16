package main

import (
	"fmt"
	"os"
	"strings"
)

type box struct {
	start int
	end   int
}

func isRowAheadClear(rowToCheck []string, boxes []box) bool {
	for _, box := range boxes {
		if rowToCheck[box.start] != "." || rowToCheck[box.end] != "." {
			return false
		}
	}
	return true
}

func isRowAheadBlocked(rowToCheck []string, boxes []box) bool {
	for _, box := range boxes {
		if rowToCheck[box.start] == "#" || rowToCheck[box.end] == "#" {
			return true
		}
	}
	return false
}

type vals struct {
	offset int
	pred   func(int, int) bool
}

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	grid := [][]string{}
  robotI, robotJ := 0, 0
	for i, line := range strings.Split(lines[0], "\n") {
		line = strings.ReplaceAll(line, "#", "##")
		line = strings.ReplaceAll(line, "O", "[]")
		line = strings.ReplaceAll(line, ".", "..")
		line = strings.ReplaceAll(line, "@", "@.")
		grid = append(grid, make([]string, len(line)))
		for j, cell := range strings.Split(line, "") {
			grid[i][j] = cell
      if cell == "@" {
        robotI, robotJ = i, j
      }
		}
	}
	moves := strings.ReplaceAll(lines[1], "\n", "")
	myMap := map[string]vals{
		"^": {-1, func(x int, y int) bool { return x < y }},
		"v": {+1, func(x int, y int) bool { return x > y }},
		"<": {-1, func(x int, y int) bool { return x < y }},
		">": {+1, func(x int, y int) bool { return x > y }},
	}
	for _, move := range strings.Split(moves, "") {
		if move == ">" || move == "<" {
			offset, pred := myMap[move].offset, myMap[move].pred
			for j := robotJ; grid[robotI][j] != "#"; j += offset {
				if grid[robotI][j] == "." {
					for ; pred(j, robotJ); j -= offset {
						grid[robotI][j] = grid[robotI][j-offset]
					}
					grid[robotI][robotJ] = "."
					robotJ += offset
					break
				}
			}
		} else if move == "^" || move == "v" {
			offset, pred := myMap[move].offset, myMap[move].pred
			rowsToMove := [][]box{{{robotJ, robotJ}}}
			for i := robotI + offset; !isRowAheadBlocked(grid[i], rowsToMove[len(rowsToMove)-1]); i += offset {
				if isRowAheadClear(grid[i], rowsToMove[len(rowsToMove)-1]) {
					for ; pred(i, robotI); i -= offset {
						boxesToMove := rowsToMove[len(rowsToMove)-1]
						rowsToMove = rowsToMove[:len(rowsToMove)-1]
						for _, box := range boxesToMove {
							grid[i][box.start], grid[i][box.end] = grid[i-offset][box.start], grid[i-offset][box.end]
							grid[i-offset][box.start], grid[i-offset][box.end] = ".", "."
						}
					}
					robotI += offset
					break
				} else {
					boxesToMove := []box{}
					set := map[box]struct{}{}
					for _, b := range rowsToMove[len(rowsToMove)-1] {
						for j := b.start - 1; j <= b.end; j++ {
							if grid[i][j] == "[" {
								set[box{j, j + 1}] = struct{}{}
							}
						}
					}
					for box := range set {
						boxesToMove = append(boxesToMove, box)
					}
					rowsToMove = append(rowsToMove, boxesToMove)
				}
			}
		}
	}
	total := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "[" {
				total += i*100 + j
			}
		}
	}
	fmt.Println(total)
}
