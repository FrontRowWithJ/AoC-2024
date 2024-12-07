package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evaluate(left, right uint64, index int, nums []uint64) bool {
	if index == len(nums)-1 {
		return left == right
	}
	concat, _ := strconv.ParseUint(fmt.Sprintf("%d%d", right, nums[index+1]), 10, 64)
	return evaluate(left, right*nums[index+1], index+1, nums) ||
		evaluate(left, right+nums[index+1], index+1, nums) ||
		evaluate(left, concat, index+1, nums)
}

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	var total uint64
	for _, line := range lines {
		ab := strings.Split(line, ": ")
		left, _ := strconv.ParseUint(ab[0], 10, 64)
		ab = strings.Split(ab[1], " ")
		nums := make([]uint64, len(ab))
		for i, s := range ab {
			nums[i], _ = strconv.ParseUint(s, 10, 64)
		}
		if evaluate(left, nums[0], 0, nums) {
			total += left
		}
	}
	fmt.Println(total)
}
