package main

import (
	"cmp"
	"fmt"
	"slices"
)

// Wouldn't have been able to do this without this post:
// https://old.reddit.com/r/adventofcode/comments/18pnycy/2023_day_24_solutions/keqf8uq/?sort=new&context=3

var axes = [3]rune{'x', 'y', 'z'}

func Part2(hailstones []Hailstone) int {
	potentialVelocities := make(map[rune][]int)
	for _, axis := range axes {
		potentialVelocities[axis] = []int{}
	}

	for i, h1 := range hailstones {
		for _, h2 := range hailstones[i+1:] {

			for _, axis := range axes {
				if h1.velocity.GetAxis(axis) == h2.velocity.GetAxis(axis) {
					validVelocities := findValidVelocities(
						int(h2.position.GetAxis(axis)-h1.position.GetAxis(axis)),
						int(h1.velocity.GetAxis(axis)),
					)

					if len(potentialVelocities[axis]) == 0 {
						potentialVelocities[axis] = validVelocities
					} else {
						potentialVelocities[axis] = sliceIntersection(potentialVelocities[axis], validVelocities)
					}
				}
			}
		}
	}

	fmt.Println("here")
	// After filtering all the intersections there should only be one valid velocity for each axis
	xVelocity := potentialVelocities['x'][0]
	yVelocity := potentialVelocities['y'][0]
	zVelocity := potentialVelocities['z'][0]

	rockVelocity := Vector3{float64(xVelocity), float64(yVelocity), float64(zVelocity)}
	// I now know its velocity, can solve its position as given two hailstones
	h1, h2 := hailstones[0], hailstones[1]
	m1 := (h1.velocity.y - rockVelocity.y) / (h1.velocity.x - rockVelocity.x)
	m2 := (h2.velocity.y - rockVelocity.y) / (h2.velocity.x - rockVelocity.x)
	c1 := h1.position.y - (m1 * h1.position.x)
	c2 := h2.position.y - (m2 * h2.position.x)
	xPos := (c2 - c1) / (m1 - m2)
	yPos := m1*xPos + c1
	time := (xPos - h1.position.x) / (h1.velocity.x - rockVelocity.x)
	zPos := h1.position.z + (h1.velocity.z-rockVelocity.z)*time

	return int(xPos + yPos + zPos)
}

// If two hailstones are going at the same speed then there's
// only a few velocities that can hit them both
// DistanceDifference % (RockVelocity-HailVelocity)
func findValidVelocities(distanceBetweenRocks int, rockVelocity int) []int {
	var validVelocities []int
	for hailV := -1000; hailV < 1000; hailV++ {
		if rockVelocity == hailV {
			continue
		}
		if distanceBetweenRocks%(rockVelocity-hailV) == 0 {
			validVelocities = append(validVelocities, hailV)
		}
	}
	return validVelocities
}

// Get elements in both a and b
func sliceIntersection[T cmp.Ordered](a, b []T) []T {
	var intersection []T
	for _, val := range a {
		_, ok := slices.BinarySearch(b, val)
		if ok {
			intersection = append(intersection, val)
		}
	}
	return intersection
}
