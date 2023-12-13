package main

import "slices"

type SpringStatus int

const (
	Operational SpringStatus = 0
	Broken      SpringStatus = 1
	Unknown     SpringStatus = 2
)

type Row struct {
	Row     []SpringStatus
	Pattern []int
}

func (row Row) GetPossibleArrangements() [][]SpringStatus {
	var unknownIndexes []int
	for i, status := range row.Row {
		i := i
		if status == Unknown {
			unknownIndexes = append(unknownIndexes, i)
		}
	}

	var arrangements [][]SpringStatus
	// Generate all possible states for the unknown cells
	for i := 0; i < 1<<len(unknownIndexes); i++ {
		possibleRow := make([]SpringStatus, len(row.Row))
		copy(possibleRow, row.Row)

		for m := 0; m < len(unknownIndexes); m++ {
			if i&(1<<m) == 0 {
				possibleRow[unknownIndexes[m]] = Operational
			} else {
				possibleRow[unknownIndexes[m]] = Broken
			}
		}

		if slices.Equal(GetPattern(possibleRow), row.Pattern) {
			arrangements = append(arrangements, possibleRow)
		}
	}

	return arrangements
}

func GetPattern(row []SpringStatus) []int {
	var patterns []int
	count := 0
	for _, status := range row {
		if status == Unknown {
			panic("unknown in arrangement")
		} else if status == Broken {
			count++
		} else if status == Operational && count > 0 {
			patterns = append(patterns, count)
			count = 0
		}
	}
	if count > 0 {
		patterns = append(patterns, count)
	}
	return patterns
}
