package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Passport has fields: byr, iyr, eyr, hgt, hcl, ecl, pid, cid
// all but cid is required
type Passport map[string]string

func vbyr(byr string) bool {
	ibyr, _ := strconv.Atoi(byr)
	return 1920 <= ibyr && ibyr <= 2002
}

func viyr(iyr string) bool {
	iiyr, _ := strconv.Atoi(iyr)
	return 2010 <= iiyr && iiyr <= 2020
}

func veyr(eyr string) bool {
	ieyr, _ := strconv.Atoi(eyr)
	return 2020 <= ieyr && ieyr <= 2030
}

func vhgt(hgt string) bool {
	ihgt, _ := strconv.Atoi(hgt[:len(hgt)-2])
	unit := hgt[len(hgt)-2:]
	switch unit {
	case "cm":
		return 150 <= ihgt && ihgt <= 193
	case "in":
		return 59 <= ihgt && ihgt <= 76
	}
	return false
}

func vhcl(hcl string) bool {
	if hcl[0] != '#' {
		return false
	}
	if len(hcl) != 7 {
		return false
	}
	for _, c := range hcl[1:] {
		if !strings.ContainsRune("abcdef0123456789", c) {
			return false
		}
	}
	return true
}

func vecl(ecl string) bool {
	switch ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

func vpid(pid string) bool {
	return len(pid) == 9
}

func isValid(passport Passport) bool {
	for _, key := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if _, prs := passport[key]; !prs {
			return false
		}
	}
	return true
}

func parseInput(filepath string) (passports []Passport) {
	fp, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	current := Passport{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, current)
			current = Passport{}
		} else {
			for _, field := range strings.Split(line, " ") {
				keyval := strings.Split(field, ":")
				current[keyval[0]] = keyval[1]
			}
		}
	}
	passports = append(passports, current)
	return
}

func main() {
	passports := parseInput("input")
	part1 := 0
	for _, passport := range passports {
		if isValid(passport) {
			part1++
		}
	}
	fmt.Println(part1)

	part2 := 0
	for _, passport := range passports {
		if isValid(passport) &&
			vbyr(passport["byr"]) &&
			vecl(passport["ecl"]) &&
			veyr(passport["eyr"]) &&
			vhcl(passport["hcl"]) &&
			vhgt(passport["hgt"]) &&
			viyr(passport["iyr"]) &&
			vpid(passport["pid"]) {
			part2++
		}
	}
	fmt.Println(part2)
}
