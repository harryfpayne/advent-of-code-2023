package part2

import (
	"github.com/harryfpayne/advent-of-code-2023/day-5/part1"
	"strconv"
	"strings"
)

func ParseMap(str string) Map {
	m := Map{}
	lines := strings.Split(str, "\n")

	for _, line := range lines[1:] {
		numbers := strings.Split(line, " ")
		if len(numbers) != 3 {
			panic("invalid map string")
		}
		destination, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		source, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(numbers[2])
		if err != nil {
			panic(err)
		}
		m.Ranges = append(m.Ranges, LookupRange{
			Source: Range{
				Start: source,
				End:   source + length,
			},
			Destination: Range{
				Start: destination,
				End:   destination + length,
			},
		})
	}

	return m
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
