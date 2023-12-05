package part2v2

import (
	"github.com/harryfpayne/advent-of-code-2023/day-5/part1"
	"math"
	"strings"
)

func Solve(puzzle string) int {
	stringMaps := strings.Split(puzzle, "\n\n")
	stringSeeds := stringMaps[0]
	seeds := ParseSeeds(stringSeeds)
	stringMaps = stringMaps[1:]

	var maps []part1.Map
	for _, m := range stringMaps {
		maps = append([]part1.Map{part1.ParseMap(m)}, maps...)
	}

	for i := 0; i < math.MaxInt; i++ {
		lookupValue := i
		for _, m := range maps {
			_lookupValue := m.GetSource(lookupValue)
			lookupValue = _lookupValue
		}

		for _, seed := range seeds {
			if seed.IsInRange(lookupValue) {
				return i
			}
		}
	}

	return -1
}
