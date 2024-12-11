package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	dict := map[uint64]uint64{}
	for _, s := range strings.Split(lines[0], " ") {
		n, _ := strconv.ParseUint(s, 10, 32)
		dict[n] += 1
	}
	var result uint64 = uint64(len(dict))
	for range 25 {
		newDict := map[uint64]uint64{}
		for number, count := range dict {
			if number == 0 {
				newDict[1] += count
			} else if s := strconv.FormatUint(number, 10); len(s)%2 == 0 {
				l, _ := strconv.ParseUint(s[:len(s)/2], 10, 64)
				r, _ := strconv.ParseUint(s[len(s)/2:], 10, 64)
				newDict[l] = newDict[l] + count
				newDict[r] = newDict[r] + count
				result += count
			} else {
				newDict[number*2024] += count
			}
		}
		dict = newDict
	}
	print(result, "\n")
}
