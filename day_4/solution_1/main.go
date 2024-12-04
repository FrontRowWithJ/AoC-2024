package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	result := 0
	for i := 0; i < len(lines)-2; i++ {
		for j := 0; j < len(lines)-2; j++ {
			word0 := string(lines[i][j]) + string(lines[i+1][j+1]) + string(lines[i+2][j+2])
			word1 := string(lines[i][j+2]) + string(lines[i+1][j+1]) + string(lines[i+2][j])
			if (word0 == "MAS" || word0 == "SAM") && (word1 == "MAS" || word1 == "SAM") {
				result +=1
			}
		}
	}
	fmt.Println(result)
}
