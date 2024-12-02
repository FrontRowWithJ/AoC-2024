package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	list0, list1 := make([]int, len(lines)), make([]int, len(lines))
	counts1 := map[int]int{}
	for i, line := range lines {
		fmt.Sscanf(line, "%d   %d", &list0[i], &list1[i])
		counts1[list1[i]]++
	}

	result := 0

	for _, n := range list0 {
		result += n * counts1[n]
	}
	fmt.Println(result)
}
