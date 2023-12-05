package part1

import (
	"math"
	"strconv"
	"strings"
)

func ParseSeeds(str string) []int {
	var seeds []int
	for _, seed := range strings.Split(strings.Replace(str, "seeds: ", " ", 1), " ") {
		if seed == "" {
			continue
		}
		s, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, s)
	}
	return seeds
}

func Solve(puzzle string) int {
	stringMaps := strings.Split(puzzle, "\n\n")
	stringSeeds := stringMaps[0]
	seeds := ParseSeeds(stringSeeds)
	stringMaps = stringMaps[1:]

	var maps []Map
	for _, m := range stringMaps {
		maps = append(maps, ParseMap(m))
	}

	minLocation := math.MaxInt
	for _, seed := range seeds {
		lookupValue := seed
		for _, m := range maps {
			_lookupValue := m.GetDestination(lookupValue)
			lookupValue = _lookupValue
		}
		if lookupValue < minLocation {
			minLocation = lookupValue
		}
	}
	return minLocation
}
