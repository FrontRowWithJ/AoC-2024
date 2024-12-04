package main

import (
	"fmt"
	"os"
	"strings"
)

func findHorizontal(lines []string) int {
	result := 0
	for _, line := range lines {
		for i := 0; i < len(line)-3; i++ {
			word := line[i : i+4]
			if word == "XMAS" || word == "SAMX" {
				result += 1
			}
		}
	}
	return result
}

func findVertical(lines []string) int {
	result := 0
	for i := 0; i < len(lines)-3; i++ {
		for j := 0; j < len(lines[0]); j++ {
			word := string(lines[i][j]) + string(lines[i+1][j]) + string(lines[i+2][j]) + string(lines[i+3][j])
			if word == "XMAS" || word == "SAMX" {
				result += 1
			}
		}
	}
	return result
}

func findDiagonalForward(lines []string) int {
	result := 0
	for i := 0; i < len(lines)-3; i++ {
		for j := 0; j < len(lines[i])-3; j++ {
			word := string(lines[i][j]) + string(lines[i+1][j+1]) + string(lines[i+2][j+2]) + string(lines[i+3][j+3])
			if word == "XMAS" || word == "SAMX" {
				result += 1
			}
		}
	}
	return result
}

func findDiagonalBackward(lines []string) int {
	result := 0
	for i := 0; i < len(lines)-3; i++ {
		for j := 3; j < len(lines[i]); j++ {
			word := string(lines[i][j]) + string(lines[i+1][j-1]) + string(lines[i+2][j-2]) + string(lines[i+3][j-3])
			if word == "XMAS" || word == "SAMX" {
				result += 1
			}
		}
	}
	return result
}

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	fmt.Println(findDiagonalBackward(lines) + findDiagonalForward(lines) + findVertical(lines) + findHorizontal(lines))
}
