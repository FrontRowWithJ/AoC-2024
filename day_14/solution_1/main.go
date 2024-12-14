package main

import (
	"fmt"
	"os"
	"strings"
)

type robot struct {
	px int
	py int
	vx int
	vy int
}

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	w, h := 101, 103
	robots := make([]robot, len(lines))
	for i, line := range lines {
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robots[i].px, &robots[i].py, &robots[i].vx, &robots[i].vy)
	}
	for index := range 10000 {
		grid := [103][101]bool{}
		for i := range robots {
			robots[i].px, robots[i].py = (robots[i].px+robots[i].vx+w)%w, (robots[i].py+robots[i].vy+h)%h
			grid[robots[i].py][robots[i].px] = true
		}
		for i := range 103 {
			for j := range 101 {
				if grid[i][j] {
					fmt.Printf("#")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("-----------------------------------------------------------------------------------------------------\n%d\n", index)
	}
}
