package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type point struct {
	i int
	j int
}

type node struct {
	score     int
	direction string
	i         int
	j         int
	path      []point
}

type scorePath struct {
	score int
	path  []point
}

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := [][]string{}
	marked := [][]map[string]int{}
	for i, line := range lines {
		grid = append(grid, strings.Split(line, ""))
		marked = append(marked, make([]map[string]int, len(grid[i])))
	}
	scorePaths := []scorePath{}
	startI, startJ := 0, 0
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "S" {
				startI, startJ = i, j
			}
			marked[i][j] = map[string]int{"^": MaxInt, "<": MaxInt, ">": MaxInt, "v": MaxInt}
		}
	}

	queue := []node{{0, ">", startI, startJ, []point{{startI, startJ}}}}
	marked[startI][startJ][">"] = 0
	for len(queue) != 0 {
		pos := queue[0]
		queue = queue[1:]
		if pos.score <= marked[pos.i][pos.j][pos.direction] {
			marked[pos.i][pos.j][pos.direction] = pos.score
		}
		if grid[pos.i][pos.j] == "E" {
			scorePaths = append(scorePaths, scorePath{pos.score, pos.path})
			continue
		}

		if pos.direction == "^" {
			if grid[pos.i-1][pos.j] != "#" && pos.score+1 <= marked[pos.i-1][pos.j]["^"] {
				marked[pos.i-1][pos.j][pos.direction] = pos.score + 1
				newPath := append(slices.Clone(pos.path), point{pos.i - 1, pos.j})
				queue = append(queue, node{pos.score + 1, "^", pos.i - 1, pos.j, newPath})
			}
			if grid[pos.i][pos.j+1] != "#" && pos.score+1001 <= marked[pos.i][pos.j+1][">"] {
				marked[pos.i][pos.j+1][pos.direction] = pos.score + 1001
				newPath := append(slices.Clone(pos.path), point{pos.i, pos.j + 1})
				queue = append(queue, node{pos.score + 1001, ">", pos.i, pos.j + 1, newPath})
			}
			if grid[pos.i][pos.j-1] != "#" && pos.score+1001 <= marked[pos.i][pos.j-1]["<"] {
				marked[pos.i][pos.j-1][pos.direction] = pos.score + 1001
				newPath := append(slices.Clone(pos.path), point{pos.i, pos.j - 1})
				queue = append(queue, node{pos.score + 1001, "<", pos.i, pos.j - 1, newPath})
			}
		} else if pos.direction == ">" {
			if grid[pos.i][pos.j+1] != "#" && pos.score+1 <= marked[pos.i][pos.j+1][">"] {
				marked[pos.i][pos.j+1][pos.direction] = pos.score + 1
				newPath := append(slices.Clone(pos.path), point{pos.i, pos.j + 1})
				queue = append(queue, node{pos.score + 1, ">", pos.i, pos.j + 1, newPath})
			}
			if grid[pos.i+1][pos.j] != "#" && pos.score+1001 <= marked[pos.i+1][pos.j]["v"] {
				marked[pos.i+1][pos.j][pos.direction] = pos.score + 1001
				newPath := append(slices.Clone(pos.path), point{pos.i + 1, pos.j})
				queue = append(queue, node{pos.score + 1001, "v", pos.i + 1, pos.j, newPath})
			}
			if grid[pos.i-1][pos.j] != "#" && pos.score+1001 <= marked[pos.i-1][pos.j]["^"] {
				marked[pos.i-1][pos.j][pos.direction] = pos.score + 1001
				newPath := append(slices.Clone(pos.path), point{pos.i - 1, pos.j})
				queue = append(queue, node{pos.score + 1001, "^", pos.i - 1, pos.j, newPath})
			}
		} else if pos.direction == "v" {
			if grid[pos.i+1][pos.j] != "#" && pos.score+1 <= marked[pos.i+1][pos.j]["v"] {
				marked[pos.i+1][pos.j][pos.direction] = pos.score + 1
				newPath := append(slices.Clone(pos.path), point{pos.i + 1, pos.j})
				queue = append(queue, node{pos.score + 1, "v", pos.i + 1, pos.j, newPath})
			}
			if grid[pos.i][pos.j-1] != "#" && pos.score+1001 <= marked[pos.i][pos.j-1]["<"] {
				marked[pos.i][pos.j-1][pos.direction] = pos.score + 1001
				newPath := append(slices.Clone(pos.path), point{pos.i, pos.j - 1})
				queue = append(queue, node{pos.score + 1001, "<", pos.i, pos.j - 1, newPath})
			}
			if grid[pos.i][pos.j+1] != "#" && pos.score+1001 <= marked[pos.i][pos.j+1][">"] {
				marked[pos.i][pos.j+1][pos.direction] = pos.score + 1001
				newPath := append(slices.Clone(pos.path), point{pos.i, pos.j + 1})
				queue = append(queue, node{pos.score + 1001, ">", pos.i, pos.j + 1, newPath})
			}
		} else if pos.direction == "<" {
			if grid[pos.i][pos.j-1] != "#" && pos.score+1 <= marked[pos.i][pos.j-1]["<"] {
				marked[pos.i][pos.j-1][pos.direction] = pos.score + 1
				newPath := append(slices.Clone(pos.path), point{pos.i, pos.j - 1})
				queue = append(queue, node{pos.score + 1, "<", pos.i, pos.j - 1, newPath})
			}
			if grid[pos.i-1][pos.j] != "#" && pos.score+1001 <= marked[pos.i-1][pos.j]["^"] {
				marked[pos.i-1][pos.j][pos.direction] = pos.score + 1001
				newPath := append(slices.Clone(pos.path), point{pos.i - 1, pos.j})
				queue = append(queue, node{pos.score + 1001, "^", pos.i - 1, pos.j, newPath})
			}
			if grid[pos.i+1][pos.j] != "#" && pos.score+1001 <= marked[pos.i+1][pos.j]["v"] {
				marked[pos.i+1][pos.j][pos.direction] = pos.score + 1001
				newPath := append(slices.Clone(pos.path), point{pos.i + 1, pos.j})
				queue = append(queue, node{pos.score + 1001, "v", pos.i + 1, pos.j, newPath})
			}
		}
	}
	slices.SortFunc(scorePaths, func(a, b scorePath) int { return a.score - b.score })
	minScore := scorePaths[0].score
	myMap := map[point]bool{}
	for _, sp := range scorePaths {
		if sp.score == minScore {
			for _, p := range sp.path {
				myMap[p] = true
			}
		}
	}
	fmt.Println(len(myMap))
}
