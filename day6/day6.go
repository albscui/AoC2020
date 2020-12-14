package main

import (
	"bufio"
	"fmt"
	"os"
)

type Group map[string]int

func main() {
	fp, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)

	// Part 1
	groups := []Group{}
	currentGroup := Group{"size": 0}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			groups = append(groups, currentGroup)
			currentGroup = Group{"size": 0}
		} else {
			for _, c := range line {
				currentGroup[string(c)]++
			}
			currentGroup["size"]++
		}
	}
	groups = append(groups, currentGroup)
	part1Count := 0
	for _, group := range groups {
		part1Count += len(group)
	}
	fmt.Println(part1Count)

	// Part 2
	part2Count := 0
	for _, g := range groups {
		for _, c := range "abcdefghijklmnopqrstuvwxyz" {
			if count, prs := g[string(c)]; prs && count == g["size"] {
				part2Count++
			}
		}
	}
	fmt.Println(part2Count)
}
