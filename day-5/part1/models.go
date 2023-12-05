package part1

import (
	"strconv"
	"strings"
)

type Map struct {
	Name         string
	LookupRanges []LookupRange
}

func (m Map) GetDestination(s int) int {
	for _, lookupRange := range m.LookupRanges {
		d := lookupRange.GetDestination(s)
		if d != -1 {
			return d
		}
	}
	return s
}

func ParseMap(str string) Map {
	m := Map{}
	lines := strings.Split(str, "\n")
	name := strings.Replace(lines[0], " map:", "", 1)
	m.Name = name

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
		m.LookupRanges = append(m.LookupRanges, LookupRange{
			SourceStart:      source,
			DestinationStart: destination,
			Length:           length,
		})
	}

	return m
}

type LookupRange struct {
	SourceStart      int
	DestinationStart int
	Length           int
}

// GetDestination Returns -1 if invalid lookup
func (l LookupRange) GetDestination(s int) int {
	if s < l.SourceStart || s > l.SourceStart+l.Length {
		return -1
	}
	return s - l.SourceStart + l.DestinationStart
}

// Added for part 2
func (m Map) GetSource(s int) int {
	for _, lookupRange := range m.LookupRanges {
		d := lookupRange.GetSource(s)
		if d != -1 {
			return d
		}
	}
	return s
}

// Added for part 2
func (l LookupRange) GetSource(s int) int {
	if s < l.DestinationStart || s > l.DestinationStart+l.Length {
		return -1
	}
	return s - l.DestinationStart + l.SourceStart
}
