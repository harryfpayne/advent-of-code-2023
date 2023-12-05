package part2

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func Solve(puzzle string) {
	stringMaps := strings.Split(puzzle, "\n\n")
	stringSeeds := stringMaps[0]
	seeds := ParseSeeds(stringSeeds)
	stringMaps = stringMaps[1:]

	var maps []Map
	for _, m := range stringMaps {
		maps = append(maps, ParseMap(m))
	}

	lowestStart := Range{1 << 61, 0}
	for _, seed := range seeds {
		ranges := []Range{seed}
		for _, m := range maps {
			var nextRanges []Range
			for _, r := range ranges {
				nextRanges = append(nextRanges, m.GetDiscontinuityPoints(r)...)
			}
			ranges = nextRanges
		}

		lowestStart = slices.MinFunc(append(ranges, lowestStart), func(a, b Range) int {
			return cmp.Compare(a.Start, b.Start)
		})
	}

	fmt.Println(lowestStart)

}
