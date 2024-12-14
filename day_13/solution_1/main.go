package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")
	totalTokens := 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		var aX, aY, bX, bY, prizeX, prizeY int
		fmt.Sscanf(s, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d",
			&aX, &aY, &bX, &bY, &prizeX, &prizeY)
		prizeX, prizeY = prizeX+10000000000000, prizeY+10000000000000
		det, detX, detY := aX*bY-bX*aY, prizeX*bY-bX*prizeY, aX*prizeY-prizeX*aY
		if det != 0 && detX%det == 0 && detY%det == 0 {
			totalTokens += detX/det*3 + detX/det
		}
	}
	fmt.Println(totalTokens)
}
