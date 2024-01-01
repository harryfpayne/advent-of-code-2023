package main

import "fmt"

type GridPoint int

const (
	Path       GridPoint = 0
	Tree       GridPoint = 1
	SlopeNorth GridPoint = 2
	SlopeEast  GridPoint = 3
	SlopeSouth GridPoint = 4
	SlopeWest  GridPoint = 5
)

func (g GridPoint) AvailableDirections(part2 bool) []int {
	if part2 {
		switch g {
		case Path, SlopeNorth, SlopeEast, SlopeSouth, SlopeWest:
			return []int{0, 1, 2, 3}
		case Tree:
			return []int{}
		default:
			panic("invalid point")
		}
	}
	switch g {
	case Path:
		return []int{0, 1, 2, 3}
	case Tree:
		return []int{}
	case SlopeNorth:
		return []int{0}
	case SlopeEast:
		return []int{1}
	case SlopeSouth:
		return []int{2}
	case SlopeWest:
		return []int{3}
	default:
		panic("invalid point")
	}
}

type Grid [][]GridPoint
type Position complex64

func (p Position) Move(direction int) Position {
	n := p
	switch direction {
	case 0:
		n -= 1i
	case 1:
		n += 1
	case 2:
		n += 1i
	case 3:
		n -= 1
	}
	return n
}

func (p Position) Y() int {
	return int(imag(p))
}
func (p Position) X() int {
	return int(real(p))
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.Y(), p.X())
}

func NewPosition(y, x int) Position {
	return Position(complex(float32(x), float32(y)))
}

func (g Grid) Get(p Position) (GridPoint, bool) {
	r := p.X()
	i := p.Y()
	if i >= len(g) || r >= len(g[0]) || i < 0 || r < 0 {
		return Path, false
	}
	return g[i][r], true
}

func (g Grid) String() string {
	var s string
	for _, row := range g {
		for _, point := range row {
			if point != Tree {
				s += "."
			} else {
				s += "#"
			}
		}
		s += "\n"
	}
	return s
}

func (g Grid) StringMap(visited map[Position]int) string {
	var s string
	for y, row := range g {
		for x, point := range row {
			var c string
			if point != Tree {
				c = "."
			} else {
				c = "#"
			}

			if val, ok := visited[NewPosition(y, x)]; ok {
				c = fmt.Sprintf("\x1b[%dm%d\x1b[0m", 41, val)
			}
			s += c
		}
		s += "\n"
	}
	return s
}
