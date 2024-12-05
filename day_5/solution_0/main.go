package main

import (
	"fmt"
	"os"
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
	page_rules := map[string]map[string]bool{}
	for _, line := range strings.Split(lines[0], "\n") {
		vals := strings.Split(line, "|")
		if len(page_rules[vals[0]]) == 0 {
			page_rules[vals[0]] = map[string]bool{}
		}
		page_rules[vals[0]][vals[1]] = true
	}
	result := 0
	for _, line := range strings.Split(lines[1], "\n") {
		result += isValid(page_rules, strings.Split(line, ","))
	}
	fmt.Println(result)
}

