package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func twoSum(nums []int, target int) (int, int, error) {
	seen := make(map[int]bool)
	for _, n := range nums {
		if _, prs := seen[target-n]; prs {
			return n, target - n, nil
		}
		seen[n] = true
	}
	return 0, 0, fmt.Errorf("Did not find twoSum")
}

func findContiguous(nums []int, target int) int {
	prevSums := map[int]int{}
	currentSum := 0
	for end, n := range nums {
		currentSum += n
		if start, prs := prevSums[currentSum-target]; prs {
			sort.Ints(nums[start+1 : end+1])
			return nums[start+1] + nums[end]
		}
		prevSums[currentSum] = end
	}
	return -1
}

func main() {
	nums := readInput("input")

	var part1 int
	for i := 25; i < len(nums); i++ {
		if _, _, err := twoSum(nums[i-25:i], nums[i]); err != nil {
			part1 = nums[i]
			break
		}
	}
	fmt.Println(part1)

	fmt.Println(findContiguous(nums, part1))
}
