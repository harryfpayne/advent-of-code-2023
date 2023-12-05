package part2

import "fmt"

type Map struct {
	Ranges []LookupRange
}

func (m Map) GetDiscontinuityPoints(r Range) []Range {
	rangesToCover := []Range{r}
	var outputRanges []Range

	lookupRangeIndex := 0

	for i := 0; i < len(rangesToCover); i++ {
		subRange := rangesToCover[i]
		//fmt.Println("------------------------------")
		if lookupRangeIndex >= len(m.Ranges) {
			// I've reached the end of lookups
			// Any remaining ranges are mapped to their value
			if len(rangesToCover) > 0 {
				outputRanges = append(outputRanges, rangesToCover[i:]...)
			}
			break
		}

		lookupRange := m.Ranges[lookupRangeIndex]
		//fmt.Println("Looking at", subRange, "and", lookupRange)
		mappedRange, missingOverlap := lookupRange.GetDiscontinuityPoints(subRange)
		if !mappedRange.IsEmpty() {
			//fmt.Println("There was an overlap!:", mappedRange)
			// I found an overlap! add it to the output
			outputRanges = append(outputRanges, mappedRange)

			if len(missingOverlap) > 0 {
				//fmt.Println("but some of it was missed:", missingOverlap)
				rangesToCover = append(rangesToCover, missingOverlap...)
				//fmt.Println("moving onto the next lookup")
				lookupRangeIndex++
			}
		} else {
			//fmt.Println("There was no overlap, using next")
			i--
			lookupRangeIndex++
		}
	}

	return outputRanges
}

type LookupRange struct {
	Source      Range
	Destination Range
}

func (l LookupRange) String() string {
	return fmt.Sprintf("%v -> %v", l.Source, l.Destination)
}

func (l LookupRange) GetDestination(s int) int {
	if s < l.Source.Start || s > l.Source.End {
		return -1
	}
	return s - l.Source.Start + l.Destination.Start
}

func (lr LookupRange) GetDiscontinuityPoints(r Range) (Range, []Range) {
	/*
		Let's say looking at range 25 - 75
		This maps 0-50   -> 200-250
		This would return [225-250], and [50-75] (because isn't covered)
	*/
	if lr.Source.Disjointed(r) {
		return Range{}, []Range{r}
	}

	if r.Contains(lr.Source) {
		// lr 10-20 | r 5-25

		var missing []Range
		if r.Start < lr.Source.Start {
			missing = append(missing, Range{
				Start: r.Start,
				End:   lr.Source.Start - 1,
			})
		}
		if r.End > lr.Source.End {
			missing = append(missing, Range{
				Start: lr.Source.End + 1,
				End:   r.End,
			})
		}

		return lr.Destination, missing
	} else if lr.Source.Contains(r) {
		// lr 10-20 | r 12-17
		destination := lr.GetDestination(r.Start)
		return Range{
			Start: destination,
			End:   destination + r.Len(),
		}, []Range{}
	} else if lr.Source.Has(r.Start) {
		// lr 10-20 | r 15-25
		destination := lr.GetDestination(r.Start)
		return Range{
				Start: destination,
				End:   lr.Destination.End,
			}, []Range{
				{
					Start: lr.Source.Start + 1,
					End:   r.End,
				},
			}
	} else if lr.Source.Has(r.End) {
		// lr 10-20 | r 5-15
		endDest := lr.GetDestination(r.End)
		return Range{
				Start: lr.Destination.Start,
				End:   endDest,
			}, []Range{
				{
					Start: r.Start,
					End:   lr.Source.Start - 1,
				},
			}
	}
	return Range{}, []Range{}
}

type Range struct {
	Start int
	End   int
}

func (r Range) Len() int {
	return r.End - r.Start
}

func (r Range) Contains(r2 Range) bool {
	return r.Start <= r2.Start && r.End >= r2.End
}

func (r Range) Has(s int) bool {
	return r.Start <= s && r.End >= s
}

func (r Range) Disjointed(r2 Range) bool {
	return r.Start > r2.End || r.End < r2.Start
}

func (r Range) IsEmpty() bool {
	return r.Start == 0 && r.End == 0
}

func (r Range) String() string {
	return fmt.Sprintf("%d-%d", r.Start, r.End)
}
