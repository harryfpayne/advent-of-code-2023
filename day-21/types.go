package main

import "fmt"

type Position complex64

func (p Position) y() int {
	return int(imag(p))
}
func (p Position) x() int {
	return int(real(p))
}

func (p Position) Adjacent() [4]Position {
	return [4]Position{
		NewPosition(p.y()-1, p.x()),
		NewPosition(p.y()+1, p.x()),
		NewPosition(p.y(), p.x()-1),
		NewPosition(p.y(), p.x()+1),
	}
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.y(), p.x())
}

func NewPosition(y, x int) Position {
	return Position(complex(float32(x), float32(y)))
}

type Grid struct {
	height int
	width  int
	g      [][]bool
}

func (g Grid) CanMove(p Position) bool {
	y := mod(p.y(), g.height)
	x := mod(p.x(), g.width)
	return !g.g[y][x]
}

func mod(a, b int) int {
	return (a%b + b) % b
}
