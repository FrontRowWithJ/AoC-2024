package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isValid(pageRules map[string]map[string]bool, nums []string) int {
	for i, n := range nums {
		if len(pageRules[n]) == 0 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if !pageRules[n][nums[j]] {
				return 0
			}
		}
	}
	n, _ := strconv.Atoi(nums[len(nums)/2])
	return n
}

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	rulesAfter := map[string]map[string]bool{}
	cmp := func(a string, b string) int {
		if len(rulesAfter[a]) != 0 && rulesAfter[a][b] {
			return 1
		}
		return -1
	}
	for _, line := range strings.Split(lines[0], "\n") {
		vals := strings.Split(line, "|")
		if rulesAfter[vals[0]] == nil {
			rulesAfter[vals[0]] = map[string]bool{}
		}
		rulesAfter[vals[0]][vals[1]] = true
	}
	result := 0
	for _, line := range strings.Split(lines[1], "\n") {
		nums := strings.Split(line, ",")
		if isValid(rulesAfter, nums) == 0 {
			slices.SortFunc(nums, cmp)
			n, _ := strconv.Atoi(nums[len(nums)/2])
			result += n
		}
	}
	fmt.Println(result)
}
