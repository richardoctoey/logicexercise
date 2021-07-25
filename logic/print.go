package logic

import "fmt"

func Print(box [][]*Box) {
	for _, v := range box {
		for _, c := range v {
			if c.ObjType == WALL {
				fmt.Print("X")
			} else if c.ObjType == GUNMAN {
				fmt.Print("P")
			} else if c.ObjType == BULLET_TRACE {
				fmt.Print("*")
			} else if c.ObjType == EMPTY {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}
