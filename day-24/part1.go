package main

func Part1(hailstones []Hailstone) int {
	intersections := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if point, ok := hailstones[i].Intersects(hailstones[j]); ok {
				if point.InBounds(min, max) &&
					hailstones[i].IsInFuture(point) &&
					hailstones[j].IsInFuture(point) {
					intersections++
				}
			}
		}
	}
	return intersections
}
