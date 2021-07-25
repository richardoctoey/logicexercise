package main

import (
	"fmt"
	"logitest/logic"
)

func main() {
	row := 4
	col := 4
	box := logic.GenerateLinkedList(row, col)
	logic.SetWall(box, [][]int{{0,0}, {0,1}, {0,2}, {2,1}, {3,0}, {3,1}, {3,2}}, row, col)
	totalGeneratedGunman := logic.SetRandomGunMan(box, row, col)
	fmt.Println("totalGeneratedGunman: ", totalGeneratedGunman)
	logic.Print(box)
}
