package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	diskLayout := []int{}
	for i, s := range strings.Split(lines[0], "") {
		for j, _ := strconv.Atoi(s); j > 0; j-- {
			diskLayout = append(diskLayout, [2]int{i / 2, -1}[i%2])
		}
	}
	for left, right := slices.Index(diskLayout, -1), len(diskLayout)-1; left < right; {
		diskLayout[left], diskLayout[right] = diskLayout[right], diskLayout[left]
		for diskLayout[right] == -1 {
			right--
		}
		left += slices.Index(diskLayout[left:], -1)
	}
	var result uint64 = 0
	for i := 0; diskLayout[i] != -1; i++ {
		result += uint64(i * diskLayout[i])
	}
	fmt.Println(result)
}
