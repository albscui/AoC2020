package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(filepath string) (output []string) {
	fp, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	return
}

func countTrees(m []string, rinc int, colinc int) (count int) {
	rows, cols := len(m), len(m[0])
	r, c := 0, 0
	for r < rows {
		if m[r][c] == '#' {
			count++
		}
		r += rinc
		c += colinc
		c %= cols
	}
	return
}

func main() {
	aocMap := readLines("input")
	part1 := countTrees(aocMap, 1, 3)
	part2 := 1
	for _, inc := range [][2]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}} {
		part2 *= countTrees(aocMap, inc[0], inc[1])
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
