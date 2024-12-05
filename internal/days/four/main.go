package four

import (
	"fmt"
)

func (ws *wordSearch) IsXmas(pos position, dir direction) bool {
	return ws.At(pos) == 'X' && ws.At(pos.step(dir, 1)) == 'M' &&
		ws.At(pos.step(dir, 2)) == 'A' && ws.At(pos.step(dir, 3)) == 'S'
}

func part1(ws *wordSearch) int {
	count := 0
	for row := 0; row < ws.Height(); row++ {
		for col := 0; col < ws.Width(); col++ {
			for _, dir := range directions {
				if ws.IsXmas(position{row: row, col: col}, dir) {
					count++
					break
				}
			}
		}
	}
	return count
}

func (ws *wordSearch) IsMasCross(pos position, dir direction) bool {
	return ws.At(pos) == 'A' &&
		ws.At(pos.step(dir, 1)) == 'S' &&
		ws.At(pos.step(dir.left(), 1)) == 'S' &&
		ws.At(pos.step(dir.right(), 1)) == 'M' &&
		ws.At(pos.step(dir.reverse(), 1)) == 'M'
}

func part2(ws *wordSearch) int {
	count := 0
	for row := 1; row < ws.Height()-1; row++ {
		for col := 1; col < ws.Width()-1; col++ {
			for _, dir := range diagonals {
				if ws.IsMasCross(position{row: row, col: col}, dir) {
					count++
				}
			}
		}
	}
	return count
}

func Main() {
	ws, err := loadWordSearch("input/4")
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(ws))
	fmt.Println(part2(ws))
}
