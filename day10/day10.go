package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var validJoltages map[int]bool
var memo map[int]int

func countArrangements(joltage int, target int) int {
	if joltage == target {
		return 1
	}
	if _, prs := memo[joltage]; !prs {
		ans := 0
		for _, i := range []int{1, 2, 3} {
			if _, prs := validJoltages[joltage+i]; prs {
				ans += countArrangements(joltage+i, target)
			}
		}
		memo[joltage] = ans
	}
	return memo[joltage]
}

func main() {
	fp, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// Parse input
	scanner := bufio.NewScanner(fp)
	joltages := []int{0}
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		joltages = append(joltages, i)
	}
	sort.Ints(joltages)
	// don't forget to add the device's joltage
	joltages = append(joltages, joltages[len(joltages)-1]+3)

	// part1
	differences := map[int]int{1: 0, 2: 0, 3: 0}
	for i := 1; i < len(joltages); i++ {
		d := joltages[i] - joltages[i-1]
		differences[d]++
	}
	fmt.Println(differences[1] * differences[3])

	// part 2
	validJoltages = map[int]bool{}
	memo = map[int]int{}
	targetJoltage := joltages[len(joltages)-1]
	for _, joltage := range joltages {
		validJoltages[joltage] = true
	}
	fmt.Println(countArrangements(0, targetJoltage))
}
