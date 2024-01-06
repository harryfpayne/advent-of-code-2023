package main

import "fmt"

type Vector3 struct {
	x, y, z float64
}

func (v Vector3) GetAxis(a rune) float64 {
	switch a {
	case 'x':
		return v.x
	case 'y':
		return v.y
	case 'z':
		return v.z
	}
	panic("invalid axis")
}

type Vector2 struct {
	x, y float64
}

func (v Vector2) InBounds(min, max float64) bool {
	return v.x >= min && v.x <= max && v.y >= min && v.y <= max
}

type Hailstone struct {
	position Vector3
	velocity Vector3
}

func (h Hailstone) String() string {
	return fmt.Sprintf("%v, %v, %v @ %v, %v, %v", h.position.x, h.position.y, h.position.z, h.velocity.x, h.velocity.y, h.velocity.z)
}

func (h Hailstone) Intersects(o Hailstone) (Vector2, bool) {
	a := Vector2{h.velocity.x, h.velocity.y}
	b := Vector2{o.velocity.x, o.velocity.y}
	c := Vector2{o.position.x - h.position.x, o.position.y - h.position.y}

	angle := CrossProduct(a, b)
	if angle == 0 { // Parallel
		return Vector2{}, false
	}

	t := CrossProduct(c, b) / angle
	return Vector2{h.position.x + h.velocity.x*t, h.position.y + h.velocity.y*t}, true
}

func CrossProduct(a, b Vector2) float64 {
	return (a.x * b.y) - (a.y * b.x)
}

func (h Hailstone) IsInFuture(point Vector2) bool {
	dx := point.x - h.position.x
	dy := point.y - h.position.y
	inXFuture := (dx > 0) == (h.velocity.x > 0)
	inYFuture := (dy > 0) == (h.velocity.y > 0)
	return inXFuture && inYFuture
}
