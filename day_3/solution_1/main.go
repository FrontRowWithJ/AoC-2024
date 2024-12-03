package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	strs := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`).FindAllString(string(input), -1)
	result := 0
	mulEnabed := true
	for _, s := range strs {
		if mulEnabed = s != "don't()" && (mulEnabed || s == "do()"); mulEnabed {
			var a, b int
			fmt.Sscanf(s, "mul(%d,%d)", &a, &b)
			result += a * b
		}
	}
	fmt.Println(result)
}
