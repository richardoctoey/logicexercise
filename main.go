package main

import (
	"flag"
	"fmt"
	"logitest/logic"
)

func main() {
	row := flag.Int("r", 4, "Rows")
	col := flag.Int("c", 4, "Columns")
	walls := flag.String("w", "{\"walls\":[[0,1]]}", "Json of Walls")
	flag.Parse()
	box := logic.GenerateLinkedList(*row, *col)
	//logic.SetWall(box, [][]int{{0,0}, {0,1}, {0,2}, {2,1}, {3,0}, {3,1}, {3,2}}, row, col)
	fmt.Println(logic.ParseJson(*walls).Walls)
	logic.SetWall(box, logic.ParseJson(*walls).Walls, *row, *col)
	totalGeneratedGunman := logic.SetRandomGunMan(box, *row, *col)
	fmt.Println("totalGeneratedGunman: ", totalGeneratedGunman)
	logic.Print(box)
}
