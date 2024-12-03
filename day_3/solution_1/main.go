package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	reg := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	res := reg.FindAll(input, -1)
	result := 0
	mulEnabed := true
	for _, mulStatement := range res {
		s := string(mulStatement)
		if s == "do()" {
			mulEnabed = true
		} else if s == "don't()" {
			mulEnabed = false
		}
		if mulEnabed {
			var a, b int
			fmt.Sscanf(string(mulStatement), "mul(%d,%d)", &a, &b)
			result += a * b
		}
	}
	fmt.Println(result)
}
