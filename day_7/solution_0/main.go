package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evaluate(left, right, index int, nums []int) bool {
	return (index == len(nums)-1 && right == left) ||
		evaluate(left, index+1, right*nums[index+1], nums) ||
		evaluate(left, index+1, right+nums[index+1], nums)
}

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	total := 0
	for _, line := range lines {
		ab := strings.Split(line, ": ")
		left, _ := strconv.Atoi(ab[0])
		ab = strings.Split(ab[1], " ")
		nums := make([]int, len(ab))
		for i, s := range ab {
			nums[i], _ = strconv.Atoi(s)
		}
		if evaluate(left, 0, nums[0], nums) {
			total += left
		}
	}
	fmt.Println(total)
}
