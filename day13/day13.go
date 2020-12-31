package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func part1(t0 int, busIDs []int) int {
	t1 := t0
	for {
		for _, busID := range busIDs {
			if busID == -1 {
				continue
			}
			if math.Mod(float64(t1), float64(busID)) == 0 {
				return busID * (t1 - t0)
			}
		}
		t1++
	}
}

func part2(t0 uint, busIDs []int) uint {
	// align with first bus
	inc := uint(busIDs[0])
	t := t0 + inc - uint(math.Mod(float64(t0), float64(inc)))
	// then keep incrementing by the product of all previous buses
	for i, busID := range busIDs[1:] {
		if busID != -1 {
			for math.Mod(float64(t+uint(i+1)), float64(busID)) != 0 {
				t += inc
			}
			inc *= uint(busID)
		}
	}
	return t
}

func main() {
	fp, err := os.Open("input")
	check(err)
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	t0, err := strconv.Atoi(scanner.Text())
	check(err)

	scanner.Scan()
	busIDs := []int{}
	for _, id := range strings.Split(scanner.Text(), ",") {
		if id != "x" {
			_id, err := strconv.Atoi(id)
			check(err)
			busIDs = append(busIDs, _id)
		} else {
			busIDs = append(busIDs, -1)
		}
	}

	fmt.Println(part1(t0, busIDs))
	fmt.Println(part2(100000000000000, busIDs))
}
