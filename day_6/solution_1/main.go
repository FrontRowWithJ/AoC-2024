package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func isInfiniteLoop(i, j int, grid [][]string) bool {
	directionI := []int{-1, 0, 1, 0}
	directionJ := []int{0, 1, 0, -1}
	directionIndex := 0
	currMovesMade := 0
	totalAllowedMoves := len(grid) * len(grid[0])
	for {
		i, j = i+directionI[directionIndex], j+directionJ[directionIndex]
		currMovesMade++
		if i == -1 || i == len(grid) || j == -1 || j == len(grid[0]) {
			return false
		}
		if grid[i][j] == "#" {
			i, j = i-directionI[directionIndex], j-directionJ[directionIndex]
			currMovesMade--
			directionIndex = (directionIndex + 1) % 4
		}
		if totalAllowedMoves == currMovesMade {
			return true
		}
	}
}

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	var guardI, guardJ int
	for guardI = 0; guardI < len(lines); guardI++ {
		guardJ = slices.Index(strings.Split(lines[guardI], ""), "^")
		if guardJ != -1 {
			break
		}
	}
	totalObstacles := 0
	for obstacleI := 0; obstacleI < len(lines); obstacleI++ {
		for obstacleJ := 0; obstacleJ < len(lines[0]); obstacleJ++ {
			if string(lines[obstacleI][obstacleJ]) == "#" {
				continue
			}
			newGrid := make([][]string, len(lines))
			for i := 0; i < len(newGrid); i++ {
				newGrid[i] = strings.Split(lines[i], "")
			}
			newGrid[obstacleI][obstacleJ] = "#"
			if isInfiniteLoop(guardI, guardJ, newGrid) {
				totalObstacles += 1
			}
		}
	}
	fmt.Println(totalObstacles)
}
