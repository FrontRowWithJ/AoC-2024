package main

import (
	"fmt"
	"os"
	"strings"
)

func find(lines []string, offSetI, offsetJ []int, iLimit, jLimit, jStart int) int {
	result := 0
	for i := 0; i < len(lines)-iLimit; i++ {
		for j := jStart; j < len(lines[i])-jLimit; j++ {
			word := string(lines[i+offSetI[0]][j+offsetJ[0]]) + string(lines[i+offSetI[1]][j+offsetJ[1]]) + string(lines[i+offSetI[2]][j+offsetJ[2]]) + string(lines[i+offSetI[3]][j+offsetJ[3]])
			if word == "XMAS" || word == "SAMX" {
				result += 1
			}
		}
	}
	return result
}

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	a := find(lines, []int{0, 0, 0, 0}, []int{0, 1, 2, 3}, 0, 3, 0) + find(lines, []int{0, 1, 2, 3}, []int{0, 0, 0, 0}, 3, 0, 0)
	b := find(lines, []int{0, 1, 2, 3}, []int{0, 1, 2, 3}, 3, 3, 0) + find(lines, []int{0, 1, 2, 3}, []int{0, -1, -2, -3}, 3, 0, 3)
	fmt.Println(a + b)
}
