package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op       string
	argument int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// execute instructions from i
func execute(instructions []instruction, i int) (int, error) {
	acc := 0
	// initiate a hash map of instruction ids
	seenInstructions := map[*instruction]bool{}
	// iterate through instructions, keep track of index of current instruction
	for i < len(instructions) {
		if _, prs := seenInstructions[&instructions[i]]; prs {
			return acc, fmt.Errorf("infinite loop found")
		}
		seenInstructions[&instructions[i]] = true
		switch instructions[i].op {
		case "nop":
			i++
		case "acc":
			acc += instructions[i].argument
			i++
		case "jmp":
			i += instructions[i].argument
		}
	}
	return acc, nil
}

func main() {
	// parse inputs into list of instructions
	fp, err := os.Open("input")
	check(err)

	scanner := bufio.NewScanner(fp)

	instructions := []instruction{}
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		argument, err := strconv.Atoi(words[1])
		check(err)
		instructions = append(instructions, instruction{words[0], argument})
	}

	// Part 1
	part1, _ := execute(instructions, 0)
	fmt.Println(part1)

	// Part 2
	for i := range instructions {
		oldOp := instructions[i].op
		if instructions[i].op == "jmp" {
			instructions[i].op = "nop"
		} else if instructions[i].op == "nop" {
			instructions[i].op = "jmp"
		}
		part2, err := execute(instructions, 0)
		if err != nil {
			instructions[i].op = oldOp
		} else {
			fmt.Println(part2)
			break
		}
	}
}
