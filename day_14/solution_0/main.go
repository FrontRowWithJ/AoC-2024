package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	w, h := 101, 103
	robots := make([][4]int, len(lines))
	for i, line := range lines {
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robots[i][0], &robots[i][1], &robots[i][2], &robots[i][3])
	}
	for range 100 {
		for i := range robots {
			robots[i][0], robots[i][1] = (robots[i][0]+robots[i][2]+w)%w, (robots[i][1]+robots[i][3]+h)%h
		}
	}
	var tl, tr, bl, br int
	for _, robot := range robots {
		if robot[0] >= 0 && robot[0] < 50 && robot[1] >= 0 && robot[1] < 51 {
			tl++
		} else if robot[0] > 50 && robot[0] < 101 && robot[1] >= 0 && robot[1] < 51 {
			tr++
		} else if robot[0] >= 0 && robot[0] < 50 && robot[1] > 51 && robot[1] < 103 {
			bl++
		} else if robot[0] > 50 && robot[0] < 101 && robot[1] > 51 && robot[1] < 103 {
			br++
		}
	}
	fmt.Println(tl * tr * bl * br)
}
