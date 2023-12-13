package main

import (
	"fmt"
)

func GetCacheKey(row []SpringStatus, pattern []int, currentPatternCount int) string {
	return fmt.Sprintf("%#v %#v %d", row, pattern, currentPatternCount)
}

var cache map[string]int = make(map[string]int)

func RecursivelyGetArrangements(row []SpringStatus, pattern []int, currentPatternCount int) int {
	cacheKey := GetCacheKey(row, pattern, currentPatternCount)
	if val, ok := cache[cacheKey]; ok {
		return val
	}

	if len(row) == 0 && len(pattern) == 0 && currentPatternCount == 0 {
		cache[cacheKey] = 1
		return 1
	}

	if len(row) == 0 && len(pattern) == 1 && currentPatternCount == pattern[0] {
		cache[cacheKey] = 1
		return 1
	}

	if len(pattern) == 0 {
		noBroken := true
		for _, status := range row {
			if status == Broken {
				noBroken = false
				break
			}
		}
		if noBroken {
			cache[cacheKey] = 1
			return 1
		}
	}

	if len(row) == 0 || len(pattern) == 0 {
		cache[cacheKey] = 0
		return 0
	}

	if currentPatternCount > pattern[0] {
		cache[cacheKey] = 0
		return 0 // Invalid to many broken to match pattern
	}

	if row[0] == Unknown {
		val := RecursivelyGetArrangements(append([]SpringStatus{Broken}, row[1:]...), pattern, currentPatternCount) +
			RecursivelyGetArrangements(append([]SpringStatus{Operational}, row[1:]...), pattern, currentPatternCount)
		cache[cacheKey] = val
		return val
	} else if row[0] == Broken {
		if currentPatternCount > pattern[0] {
			cache[cacheKey] = 0
			return 0 // Invalid to many broken to match pattern
		}
		val := RecursivelyGetArrangements(row[1:], pattern, currentPatternCount+1)
		cache[cacheKey] = val
		return val
	} else {
		if currentPatternCount > 0 && currentPatternCount != pattern[0] {
			cache[cacheKey] = 0
			return 0 // Invalid not correct number of broken
		}
		if currentPatternCount > 0 {
			val := RecursivelyGetArrangements(row[1:], pattern[1:], 0)
			cache[cacheKey] = val
			return val
		}
		val := RecursivelyGetArrangements(row[1:], pattern, 0)
		cache[cacheKey] = val
		return val
	}
}
