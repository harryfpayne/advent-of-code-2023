package main

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func GetLessThanCount(a int, arr []int) int {
	count := 0
	for _, b := range arr {
		if a > b {
			count++
		}
	}
	return count
}
