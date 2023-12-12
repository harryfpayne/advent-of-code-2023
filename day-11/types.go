package main

type Space bool

type Point [2]int

type Map [][]Space

var ExpansionFactor = 2

func (m Map) GetPointsAfterSpaceTimeExpansion() []Point {
	var rowsForExpansion []int
	var columnsForExpansion []int

	for i := 0; i < len(m); i++ {
		rowEmpty := true
		columnEmpty := true
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] {
				rowEmpty = false
			}
			if m[j][i] {
				columnEmpty = false
			}
		}
		if rowEmpty {
			rowsForExpansion = append(rowsForExpansion, i)
		}
		if columnEmpty {
			columnsForExpansion = append(columnsForExpansion, i)
		}
	}

	var points []Point
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] {
				// I'm looking at a star
				// How many shifts have there been in my row?
				xExpansion := GetLessThanCount(x, columnsForExpansion) * (ExpansionFactor - 1)
				// How many shifts have there been in my column?
				yExpansion := GetLessThanCount(y, rowsForExpansion) * (ExpansionFactor - 1)
				points = append(points, Point{y + yExpansion, x + xExpansion})
			}
		}
	}

	return points
}
