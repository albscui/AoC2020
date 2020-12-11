package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Password ..
type Password struct {
	password string
	target   byte
	min      int
	max      int
}

func readInput(filepath string) (output []Password) {
	fp, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		// 17-18 k: kkkkkkkkwkkkkkkkbmk
		line := strings.Split(scanner.Text(), " ")
		constraints, target, password := strings.Split(line[0], "-"), line[1][0], line[2]
		min, _ := strconv.Atoi(constraints[0])
		max, _ := strconv.Atoi(constraints[1])
		output = append(output, Password{password, target, min, max})
	}
	return
}

func validatePassword(password string, target byte, min int, max int) bool {
	count := 0
	for _, c := range password {
		if byte(c) == target {
			count++
		}
	}
	return min <= count && count <= max
}

func validatePassword2(password string, target byte, idx1 int, idx2 int) bool {
	a := password[idx1-1] == target
	b := password[idx2-1] == target
	return (a || b) && !(a && b)
}

func main() {

	// fmt.Println(validatePassword("gqdrspndrpsrjfjx", 'r', 1, 2))
	part1ans, part2ans := 0, 0
	for _, pwd := range readInput("input") {
		if validatePassword(pwd.password, pwd.target, pwd.min, pwd.max) {
			part1ans++
		}
		if validatePassword2(pwd.password, pwd.target, pwd.min, pwd.max) {
			part2ans++
		}
	}
	fmt.Println(part1ans)
	fmt.Println(part2ans)
}
