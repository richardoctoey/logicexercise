package main

import (
	"fmt"
	"logitest/logic"
)

func main() {
	box := logic.GenerateLinkedList(4, 4)
	for _, v := range box {
		for _, c := range v {
			var curr, next, prev, top, bot int
			curr = c.Value
			if c.Next != nil {
				next = c.Next.Value
			}
			if c.Prev != nil {
				prev = c.Prev.Value
			}
			if c.Top != nil {
				top = c.Top.Value
			}
			if c.Down != nil {
				bot = c.Down.Value
			}

			fmt.Print(fmt.Sprintf("curr: %d, next: %d, prev: %d, top: %d, bot: %d | ", curr, next, prev, top, bot))
		}
		fmt.Println()
		fmt.Println("Next Line")
	}
}
