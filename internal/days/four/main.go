package four

import (
	"bufio"
	"fmt"
	"os"
)

type wordSearch struct {
	grid [][]rune
}

type position struct {
	row int
	col int
}

type direction struct {
	x int
	y int
}

var diagonals = []direction{
	{x: 1, y: 1},
	{x: -1, y: -1},
	{x: 1, y: -1},
	{x: -1, y: 1},
}

var directions = []direction{
	{x: 0, y: 1}, // orthogonal
	{x: 1, y: 0},
	{x: 0, y: -1},
	{x: -1, y: 0},
	{x: 1, y: 1}, // diagonal
	{x: -1, y: -1},
	{x: 1, y: -1},
	{x: -1, y: 1},
}

func (p position) step(d direction, steps int) position {
	return position{row: p.row + d.y*steps, col: p.col + d.x*steps}
}

func (d direction) reverse() direction {
	return direction{x: -d.x, y: -d.y}
}

func (d direction) left() direction {
	return direction{x: -d.y, y: d.x}
}

func (d direction) right() direction {
	return direction{x: d.y, y: -d.x}
}

func loadWordSearch(path string) (*wordSearch, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	grid := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &wordSearch{grid: grid}, nil
}

func (ws *wordSearch) Width() int {
	return len(ws.grid[0])
}

func (ws *wordSearch) Height() int {
	return len(ws.grid)
}

func (ws *wordSearch) At(p position) rune {
	if p.row < 0 || p.row >= ws.Height() || p.col < 0 || p.col >= ws.Width() {
		return 0
	}
	return ws.grid[p.row][p.col]
}

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
