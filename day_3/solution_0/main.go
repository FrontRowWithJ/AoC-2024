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
	
	reg:= regexp.MustCompile(`mul\(\d+,\d+\)`)
	res := reg.FindAll(input, -1)
	result := 0
	for _, mulStatement := range res {
		var a, b int
		fmt.Sscanf(string(mulStatement), "mul(%d,%d)", &a, &b)
		result += a * b
	}
	fmt.Println(result)
}