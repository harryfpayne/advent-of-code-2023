package main

import "fmt"

type Cube struct {
	x int
	y int
	z int
}

func (c *Cube) Below() Cube {
	return Cube{
		x: c.x,
		y: c.y,
		z: c.z - 1,
	}
}

func (c *Cube) Above() Cube {
	return Cube{
		x: c.x,
		y: c.y,
		z: c.z + 1,
	}
}

type Brick []Cube

func Print(bricks []Brick, dir rune) {
	maxCube := Cube{}
	grid := make(map[Cube]int)
	for i, brick := range bricks {
		for _, cube := range brick {
			if cube.z > maxCube.z {
				maxCube.z = cube.z
			}
			if cube.y > maxCube.y {
				maxCube.y = cube.y
			}
			if cube.x > maxCube.x {
				maxCube.x = cube.x
			}
			grid[cube] = i
		}
	}

	for z := maxCube.z; z >= 0; z-- {
		for a := 0; a <= maxCube.x; a++ {
			iInThisRow := -1
			for b := 0; b <= maxCube.y; b++ {
				c := Cube{a, b, z}
				if dir == 'y' {
					c = Cube{b, a, z}
				}
				if val, ok := grid[c]; ok {
					iInThisRow = val
				}
			}
			if iInThisRow != -1 {
				fmt.Print(iInThisRow)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func PrintMatrix(supportMatrix [][]bool) {
	println("  0 1 2 3 4 5 6")
	for y, row := range supportMatrix {
		print(y, "|")
		for _, b := range row {
			if b {
				print("x ")
			} else {
				print(". ")
			}
		}
		println()
	}
}

func u(i int, e error) int {
	if e != nil {
		panic(e)
	}
	return i
}

func GeneratePositionMap(bricks []Brick) map[Cube]int {
	grid := make(map[Cube]int)
	for i, brick := range bricks {
		for _, cube := range brick {
			grid[cube] = i
		}
	}
	return grid
}
