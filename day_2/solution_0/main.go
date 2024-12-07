package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isIncreasingOrDecreasing(arr []int) int {
	isIncreasing := arr[0] < arr[1]
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		absDiff := abs(diff)
		if (isIncreasing != (diff > 0)) || absDiff < 1 || absDiff > 3 {
			return 0
		}
	}
	return 1
}

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	total := 0
	for _, line := range lines {
		ss := strings.Fields(line)
		arr := make([]int, len(ss))
		for i, s := range ss {
			arr[i], _ = strconv.Atoi(s)
		}
		total += isIncreasingOrDecreasing(arr)
	}
	fmt.Println(total)
}
