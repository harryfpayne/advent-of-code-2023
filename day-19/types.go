package main

type Rule struct {
	key         uint8
	comparator  bool // false = >, true = <
	value       int
	destination string
}

type Instruction struct {
	name     string
	rules    []Rule
	fallback string
}

type Part struct {
	x int
	m int
	a int
	s int
}

func (p Part) Sum() int {
	return p.x + p.a + p.m + p.s
}

func (p Part) MeetsRequirement(r Rule) bool {
	val := 0
	switch r.key {
	case 'x':
		val = p.x
	case 'm':
		val = p.m
	case 'a':
		val = p.a
	case 's':
		val = p.s
	default:
		panic("invalid key")
	}

	if r.comparator {
		return val < r.value
	} else {
		return val > r.value
	}
}
