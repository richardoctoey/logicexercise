package logic

import (
	"fmt"
	"go/types"
	"unsafe"
)

type ObjectType int
const (
	EMPTY ObjectType = 1
	GUNMAN ObjectType = 2
	WALL ObjectType = 3
)


type Box struct {
	Next *Box
	Prev *Box
	Top *Box
	Down *Box
	ObjType ObjectType
}

func isTopLeftCorner(r int, c int, row int, col int) bool {
	return r == 0 && c == 0
}

func isTopRightCorner(r int, c int, row int, col int) bool {
	return r == 0 && c == col
}

func isBottomLeftCorner(r int, c int, row int, col int) bool {
	return r == row && c == 0
}

func isBottomRightCorner(r int, c int, row int, col int) bool {
	return r == row && c == col
}

func isTopLineExceptCorner(r int, c int, row int, col int) bool {
	return r == 0 && c != 0 && c != col
}

func isBottomLineExceptCorner(r int, c int, row int, col int) bool {
	return r == row && c != 0 && c != col
}

func isLeftLineExceptCorner(r int, c int, row int, col int) bool {
	return c == 0 && r != 0 && r != row
}

func isRightLineExceptCorner(r int, c int, row int, col int) bool {
	return c == col && r != 0 && r != row
}

func GenerateLinkedList(row int, col int) {
	totalBox := row * col
	row := row - 1
	col := col - 1
	arrBox := [][]*Box{}
	for c := 0; c <= col; c++ {
		for r := 0; r <= row; r++ {
			var nextObj, prevObj, topObj, downObj *Box
			index := c * r
			if isTopLeftCorner(r, c, row, col) {
				prevObj = nil
			} else if isTopRightCorner(r, c, row, col) {
				nextObj = nil
				prevObj = arrBox[r][c-1]
			} else if isBottomLeftCorner(r, c, row, col) {
				prevObj = nil
			} else if isBottomRightCorner(r, c, row, col) {
				nextObj = nil
			} else if isTopLineExceptCorner(r, c, row, col)  {
				prevObj = arrBox[r][c-1]
				prevObj.Next = arrBox[r][c]
			}
		}
	}
}
