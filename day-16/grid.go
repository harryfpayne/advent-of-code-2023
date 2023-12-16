package main

import (
	"fmt"
	"time"
)

type Position [2]int

type Beam struct {
	position Position
	facing   int //N=0, E=1, S=2, E=3
}

type Grid struct {
	layout     [][]int32
	beams      []Beam
	beamPoints map[Beam]struct{}
}

const log = false

func (grid Grid) GetCoverageFrom(beam Beam) int {
	grid.beams = grid.TransformBeam(beam)
	grid.beamPoints[beam] = struct{}{}

	for len(grid.beams) > 0 {
		if log {
			fmt.Printf("\033[0;0H")
			fmt.Println(grid)
			time.Sleep(time.Second / 20)
		}
		grid.MoveBeams()
	}

	uniquePoints := make(map[Position]struct{})
	for beam, _ := range grid.beamPoints {
		uniquePoints[beam.position] = struct{}{}
	}

	a := len(uniquePoints)
	return a
}

func (g *Grid) MoveBeams() {
	var nextBeams []Beam

	for _, beam := range g.beams {
		beam := beam
		beam.Move()

		next := g.TransformBeam(beam)
		if len(next) == 0 { // This beam is dead
			continue
		}
		if _, ok := g.beamPoints[beam]; ok { // If I've already been here end this beam
			continue
		} else {
			g.beamPoints[beam] = struct{}{}
		}
		nextBeams = append(nextBeams, next...)
	}

	g.beams = nextBeams
}

func (g *Grid) TransformBeam(beam Beam) (nextBeams []Beam) {
	movedOnto, ok := g.Get(beam.position)
	if !ok { // If can't move end this beam
		return
	}
	switch movedOnto {
	case '/':
		switch beam.facing {
		case 0:
			beam.facing = 1
		case 1:
			beam.facing = 0
		case 2:
			beam.facing = 3
		case 3:
			beam.facing = 2
		}
		nextBeams = append(nextBeams, beam)
	case '\\':
		switch beam.facing {
		case 0:
			beam.facing = 3
		case 1:
			beam.facing = 2
		case 2:
			beam.facing = 1
		case 3:
			beam.facing = 0
		}
		nextBeams = append(nextBeams, beam)
	case '|':
		if beam.facing == 1 || beam.facing == 3 {
			nextBeams = append(nextBeams,
				Beam{beam.position, 0},
				Beam{beam.position, 2},
			)
		} else {
			nextBeams = append(nextBeams, beam)
		}
	case '-':
		if beam.facing == 0 || beam.facing == 2 {
			nextBeams = append(nextBeams,
				Beam{beam.position, 1},
				Beam{beam.position, 3},
			)
		} else {
			nextBeams = append(nextBeams, beam)
		}
	default:
		nextBeams = append(nextBeams, beam)
	}
	return
}

func (g *Grid) Get(p Position) (int32, bool) {
	if p[0] >= len(g.layout) || p[1] >= len(g.layout[0]) || p[0] < 0 || p[1] < 0 {
		return 0, false
	}
	return g.layout[p[0]][p[1]], true
}

var PossibleDirections = [...]int{0, 1, 2, 3}

func (g Grid) String() (s string) {
	for y, line := range g.layout {
		for x, p := range line {
			if p != '.' {
				s += string(p)
				continue
			}

			found := false
			for _, d := range PossibleDirections {
				if _, ok := g.beamPoints[Beam{Position{y, x}, d}]; ok {
					s += "#"
					found = true
					break
				}
			}
			if !found {
				s += "."
			}
		}
		s += "\n"
	}
	return
}

func (b *Beam) Move() {
	switch b.facing {
	case 0:
		b.position = [2]int{b.position[0] - 1, b.position[1]}
	case 1:
		b.position = [2]int{b.position[0], b.position[1] + 1}
	case 2:
		b.position = [2]int{b.position[0] + 1, b.position[1]}
	case 3:
		b.position = [2]int{b.position[0], b.position[1] - 1}
	default:
		panic("invalid direction")
	}
}
