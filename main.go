package main

import (
	"math/rand"
	"syscall/js"
	"time"
)

const (
	ROW = 40
	COL = 100
)

var cells = [ROW][COL]int{}
var tmpCells = [ROW][COL]int{}

func main() {
	initCells()

	document := js.Global().Get("document")
	pre := document.Call("getElementById", "life-game")

	for {
		time.Sleep(200 * time.Millisecond)
		updateAll()
		str := convertCellsToString()
		pre.Set("textContent", str)
	}
}

func initCells() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			randNum := rand.Intn(2)
			if randNum == 0 {
				cells[i][j] = 1
			} else {
				cells[i][j] = 0
			}
		}
	}
}

func updateAll() {
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			tmpCells[i][j] = updateCell(i, j)
		}
	}
	cells = tmpCells
}

func updateCell(rowIdx int, colIdx int) int {
	cnt := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			tx := rowIdx + dx
			ty := colIdx + dy
			if (dx == 0 && dy == 0) || !(0 <= tx && tx < ROW && 0 <= ty && ty < COL) {
				continue
			}
			cnt += cells[tx][ty]

		}
	}
	switch cnt {
	case 2:
		return cells[rowIdx][colIdx]
	case 3:
		return 1
	default:
		return 0
	}
}

func convertCellsToString() string {
	str := ""
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			var now string
			if cells[i][j] == 0 {
				now = "□"
			} else {
				now = "■"
			}
			str = str + now
		}
		str = str + "\n"
	}
	return str
}
