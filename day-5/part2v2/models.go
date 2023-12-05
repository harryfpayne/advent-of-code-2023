package part2v2

import (
	"github.com/harryfpayne/advent-of-code-2023/day-5/part1"
)

type Range struct {
	Start int
	End   int
}

func (r Range) IsInRange(n int) bool {
	return r.Start <= n && r.End >= n
}

func ParseSeeds(str string) []Range {
	singleSeeds := part1.ParseSeeds(str)

	var seeds []Range
	for i := 0; i < len(singleSeeds); i += 2 {
		seeds = append(seeds, Range{
			Start: singleSeeds[i],
			End:   singleSeeds[i] + singleSeeds[i+1],
		})
	}

	return seeds
}
