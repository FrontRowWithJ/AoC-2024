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

func isIncreasingOrDecreasing(arr []int) bool {
	isIncreasing := arr[0] < arr[1]
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		absDiff := abs(diff)

		if (isIncreasing != (diff > 0)) || absDiff < 1 || absDiff > 3 {
			return false
		}
	}
	return true
}

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	total := 0
	for _, line := range lines {
		var arr []int
		for _, s := range strings.Fields(line) {
			n, _ := strconv.ParseInt(s, 10, 32)
			arr = append(arr, int(n))
		}
		if isIncreasingOrDecreasing(arr) {
			total += 1
		}
	}
	fmt.Println(total)
}
