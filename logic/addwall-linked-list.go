package logic

func bentrok(mytype ObjectType) {
	if mytype != EMPTY {
		panic("Ooops bentrok!!")
	}
}

func setBulletTrace(currBox *Box) {
	selectedObj := currBox
	for selectedObj.Next != nil && selectedObj.Next.ObjType != WALL {
		selectedObj = selectedObj.Next
		selectedObj.ObjType = BULLET_TRACE
	}
	selectedObj = currBox
	for selectedObj.Prev != nil && selectedObj.Prev.ObjType != WALL {
		selectedObj = selectedObj.Prev
		selectedObj.ObjType = BULLET_TRACE
	}
	selectedObj = currBox
	for selectedObj.Top != nil && selectedObj.Top.ObjType != WALL {
		selectedObj = selectedObj.Top
		selectedObj.ObjType = BULLET_TRACE
	}
	selectedObj = currBox
	for selectedObj.Down != nil && selectedObj.Down.ObjType != WALL {
		selectedObj = selectedObj.Down
		selectedObj.ObjType = BULLET_TRACE
	}
}

func SetWall(arrBox [][]*Box, walls [][]int, totalRow int, totalCol int) {
	for row, _ := range walls {
		r := walls[row][0]
		c := walls[row][1]
		var curr *Box
		curr = arrBox[r][c]
		curr.ObjType = WALL
	}
}

func SetGunMan(arrBox [][]*Box, gunmans [][]int, totalRow int, totalCol int) {
	for row, _ := range gunmans {
		r := gunmans[row][0]
		c := gunmans[row][1]
		var curr *Box
		curr = arrBox[r][c]
		bentrok(curr.ObjType)
		curr.ObjType = GUNMAN
		setBulletTrace(curr)
	}
}

func SetRandomGunMan(arrBox [][]*Box, totalRow int, totalCol int) int {
	totalGenerateGunMan := 0
	for row, _ := range arrBox {
		for col, _ := range arrBox[row] {
			if arrBox[row][col].ObjType == EMPTY {
				SetGunMan(arrBox, [][]int{{row, col}}, totalRow, totalCol)
				totalGenerateGunMan += 1
			}
		}
	}
	return totalGenerateGunMan
}
