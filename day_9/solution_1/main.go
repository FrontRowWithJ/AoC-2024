package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	diskLayout := []int{}
	usedBlocks, freeBlocks := [][2]int{}, [][2]int{}
	ptrs := [2]*[][2]int{&usedBlocks, &freeBlocks}
	for i, s := range strings.Split(lines[0], "") {
		spaceLen, _ := strconv.Atoi(s)
		*ptrs[i%2] = append(*ptrs[i%2], [2]int{len(diskLayout), spaceLen})
		for j, _ := strconv.Atoi(s); j > 0; j-- {
			diskLayout = append(diskLayout, [2]int{i / 2, 0}[i%2])
		}
	}
	for blockID := len(usedBlocks) - 1; blockID >= 0; blockID-- {
		blockPosition, blockLength := usedBlocks[blockID][0], usedBlocks[blockID][1]
		for i := range freeBlocks {
			position, length := freeBlocks[i][0], freeBlocks[i][1]
			if length >= blockLength && position < blockPosition {
				for j := 0; j < blockLength; j++ {
					diskLayout[position+j], diskLayout[blockPosition+j] = diskLayout[blockPosition+j], diskLayout[position+j]
				}
				freeBlocks[i][0], freeBlocks[i][1] = freeBlocks[i][0]+blockLength, freeBlocks[i][1]-blockLength
				break
			}
		}
	}
	var result uint64 = 0
	for i, id := range diskLayout {
		result += uint64(i * id)
	}
	fmt.Println(result)
}
