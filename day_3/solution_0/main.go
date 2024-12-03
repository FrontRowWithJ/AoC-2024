package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	res := regexp.MustCompile(`mul\(\d+,\d+\)`).FindAllString(string(input), -1)
	result := 0
	for _, mulStatement := range res {
		var a, b int
		fmt.Sscanf(mulStatement, "mul(%d,%d)", &a, &b)
		result += a * b
	}
	fmt.Println(result)
}
