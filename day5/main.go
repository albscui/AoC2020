package main

import (
	"bufio"
	"fmt"
	"os"
)

func flightID(boardingPass string) int {
	lo, hi := 0, 127
	var m int
	for _, c := range boardingPass[:7] {
		m = lo + (hi-lo)/2
		switch c {
		case 'F':
			hi = m
		case 'B':
			lo = m + 1
		}
	}
	row := lo + (hi-lo)/2

	lo, hi = 0, 7
	for _, c := range boardingPass[7:] {
		m = lo + (hi-lo)/2
		switch c {
		case 'R':
			lo = m + 1
		case 'L':
			hi = m
		}
	}
	col := lo + (hi-lo)/2
	return row*8 + col
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Given an array containing max - min numbers in range [min, max], find the
// only missing number.
func findMissing(nums []int, _min int) int {
	// Extend the array by 1 to cover entire range
	nums = append(nums, -1)
	for _, n := range nums {
		j := abs(n) - _min
		if j >= 0 && nums[j] > 0 {
			nums[j] = -nums[j]
		}
	}
	// The missing number must correspond to the index at which the number didn't
	// get turned into a negative.
	for i, n := range nums {
		if n > 0 {
			return i + _min
		}
	}
	return -1
}

func main() {
	fp, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)

	// Part 1
	maxID, minID := 0, 1<<63-1
	ids := []int{}
	for scanner.Scan() {
		id := flightID(scanner.Text())
		ids = append(ids, id)
		maxID = max(maxID, id)
		minID = min(minID, id)
	}
	fmt.Println(maxID)

	// Part 2
	// Find the mising id without sorting
	fmt.Println(findMissing(ids, minID))
}
