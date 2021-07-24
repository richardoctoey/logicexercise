package logic

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
	Value int
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

func GenerateLinkedList(row int, col int) [][]*Box {
	row = row - 1
	col = col - 1
	arrBox := [][]*Box{}
	val := 0
	for r := 0; r <= row; r++ {
		row := []*Box{}
		for c := 0; c <= col; c++ {
			row = append(row, &Box{Value: val})
			val+=1
		}
		arrBox = append(arrBox, row)
	}
	for r := 0; r <= row; r++ {
		for c := 0; c <= col; c++ {
			var curObj *Box
			curObj = arrBox[r][c]
			if isTopLeftCorner(r, c, row, col) {
				curObj.Next = arrBox[r][c+1]
				curObj.Down = arrBox[r+1][c]
			} else if isTopRightCorner(r, c, row, col) {
				curObj.Prev = arrBox[r][c-1]
				curObj.Down = arrBox[r+1][c]
			} else if isBottomLeftCorner(r, c, row, col) {
				curObj.Next = arrBox[r][c+1]
				curObj.Top = arrBox[r-1][c]
			} else if isBottomRightCorner(r, c, row, col) {
				curObj.Prev = arrBox[r][c-1]
				curObj.Top = arrBox[r-1][c]
			} else if isTopLineExceptCorner(r, c, row, col)  {
				curObj.Prev = arrBox[r][c-1]
				curObj.Next = arrBox[r][c+1]
				curObj.Down = arrBox[r+1][c]
			} else if isBottomLineExceptCorner(r, c, row, col) {
				curObj.Prev = arrBox[r][c-1]
				curObj.Next = arrBox[r][c+1]
				curObj.Top = arrBox[r-1][c]
			} else if isLeftLineExceptCorner(r, c, row, col) {
				curObj.Next = arrBox[r][c+1]
				curObj.Top = arrBox[r-1][c]
				curObj.Down = arrBox[r+1][c]
			} else if isRightLineExceptCorner(r, c, row, col) {
				curObj.Prev = arrBox[r][c-1]
				curObj.Top = arrBox[r-1][c]
				curObj.Down = arrBox[r+1][c]
			} else {
				curObj.Prev = arrBox[r][c-1]
				curObj.Next = arrBox[r][c+1]
				curObj.Top = arrBox[r-1][c]
				curObj.Down = arrBox[r+1][c]
			}
		}
	}
	return arrBox
}
