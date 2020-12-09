package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(filepath string) (output []int) {
	fp, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		output = append(output, i)
	}
	return
}

func twoSum(nums []int, target int) (int, int) {
	seen := make(map[int]bool)
	for _, n := range nums {
		if _, prs := seen[target-n]; prs {
			return n, target - n
		}
		seen[n] = true
	}
	return 0, 0
}

// Given a list of numbers and a target, find three numbers x1, x2, x3 such that x1 + x2 + x3 = target
// also x1 x2 x3 must be distinct elements in the list.
func threeSum(nums []int, target int) (int, int, int) {
	for i, x1 := range nums {
		x2, x3 := twoSum(nums[i+1:], target-x1)
		if x1+x2+x3 == target {
			return x1, x2, x3
		}
	}
	return 0, 0, 0
}

func main() {
	nums := readInput("input")
	x1, x2 := twoSum(nums, 2020)
	fmt.Printf("Part 1: %d x %d = %d\n", x1, x2, x1*x2)
	x1, x2, x3 := threeSum(nums, 2020)
	fmt.Printf("Part 2: %d x %d x %d = %d\n", x1, x2, x3, x1*x2*x3)
}
