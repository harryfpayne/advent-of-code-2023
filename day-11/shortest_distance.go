package main

func GetShortestPairDistance(points []Point) int {
	distanceSum := 0
	for a := 0; a < len(points); a++ {
		for b := 0; b < len(points); b++ {
			if b >= a {
				continue
			}

			distanceSum += Distance(points[a], points[b])
		}
	}

	return distanceSum
}

func Distance(a, b Point) int {
	yDist := b[0] - a[0]
	xDist := b[1] - a[1]
	return Abs(yDist) + Abs(xDist)
}
