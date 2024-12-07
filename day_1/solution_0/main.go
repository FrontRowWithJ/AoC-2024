package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	list0, list1 := make([]int, len(lines)), make([]int, len(lines))
	for i, line := range lines {
		fmt.Sscanf(line, "%d   %d", &list0[i], &list1[i])
	}
	slices.Sort(list0)
	slices.Sort(list1)
	result := 0
	for i := 0; i < len(list0); i++ {
		result += abs(list0[i] - list1[i])
	}
	fmt.Println(result)
}
