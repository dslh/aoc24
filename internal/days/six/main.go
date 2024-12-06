package six

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type position struct {
	x int
	y int
}

func (p position) right() position {
	return position{-p.y, p.x}
}
func (p position) add(o position) position {
	return position{p.x + o.x, p.y + o.y}
}

type grid struct {
	m      map[position]struct{}
	width  int
	height int
}

func newGrid() *grid {
	return &grid{m: make(map[position]struct{})}
}

func (g *grid) set(p position) {
	g.m[p] = struct{}{}
	if p.x > g.width {
		g.width = p.x
	}
	if p.y > g.height {
		g.height = p.y
	}
}
func (g *grid) get(p position) bool {
	_, ok := g.m[p]
	return ok
}
func (g *grid) count() int {
	return len(g.m)
}
func (g *grid) within(p position) bool {
	return p.x >= 0 && p.x <= g.width && p.y >= 0 && p.y <= g.height
}

func readInput() (*grid, position, error) {
	file, err := os.Open("input/6")
	if err != nil {
		return nil, position{}, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	g := newGrid()
	p := position{0, 0}
	y := 0
	x := 0
	for {
		ch, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, position{}, err
		}

		switch ch {
		case '.':
			x++
		case '#':
			g.set(position{x, y})
			x++
		case '\n':
			y++
			x = 0
		case '^':
			p = position{x, y}
			x++
		default:
			return nil, position{}, fmt.Errorf("unknown character: %c", ch)
		}
	}

	return g, p, nil
}

func walk(g *grid, p position, d position) *grid {
	visited := newGrid()
	for {
		visited.set(p)
		next := p.add(d)
		if !g.within(next) {
			break
		}
		if g.get(next) {
			d = d.right()
		} else {
			p = next
		}
	}
	return visited
}

func print(g *grid, v *grid) {
	for y := 0; y <= g.height; y++ {
		for x := 0; x <= g.width; x++ {
			if v.get(position{x, y}) {
				fmt.Print("X")
			} else if g.get(position{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Main() {
	grid, start, err := readInput()
	if err != nil {
		panic(err)
	}

	visited := walk(grid, start, position{0, -1})
	print(grid, visited)
	fmt.Println(visited.count())
}
