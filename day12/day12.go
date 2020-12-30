package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	N = iota
	S
	W
	E
	L
	R
	F
)

type Mover interface {
	N(int)
	S(int)
	W(int)
	E(int)
	L(int)
	R(int)
	F(int)
}

type Ship struct {
	x, y      int
	direction int
}

func (s *Ship) N(n int) {
	s.move(N, n)
}
func (s *Ship) S(n int) {
	s.move(S, n)
}
func (s *Ship) E(n int) {
	s.move(E, n)
}
func (s *Ship) W(n int) {
	s.move(W, n)
}
func (s *Ship) L(n int) {
	s.rotate([]int{E, N, W, S}, n)
}
func (s *Ship) R(n int) {
	s.rotate([]int{E, S, W, N}, n)
}
func (s *Ship) F(n int) {
	s.move(s.direction, n)
}

func (s *Ship) move(direction, n int) {
	switch direction {
	case N:
		s.y += n
	case S:
		s.y -= n
	case E:
		s.x += n
	case W:
		s.x -= n
	}
}

func (s *Ship) rotate(directions []int, angle int) {
	// angles: 0, 90, 180, 270, 360
	for i, d := range directions {
		if s.direction == d {
			_idx := int(math.Mod(float64(i+angle/90), 4))
			s.direction = directions[_idx]
			break
		}
	}
}

func (s Ship) taxiCabDist() float64 {
	return math.Abs(float64(s.x)) + math.Abs(float64(s.y))
}

type Waypoint struct {
	x, y int // relative to a ship as the origin
}

func (w *Waypoint) move(direction, n int) {
	switch direction {
	case N:
		w.y += n
	case S:
		w.y -= n
	case W:
		w.x -= n
	case E:
		w.x += n
	}
}

// rotate counter clockwise by angle a
// _x = x * cos(a) - y * sin(a)
// _y = x * sin(a) + y * cos(a)
func (w *Waypoint) ccw(a int) {
	for i := 0; i < a/90; i++ {
		w.x, w.y = -w.y, w.x
	}
}

// rotate clockwise by angle a
// _x = x * cos(a) + y * sin(a)
// _y = x * -sin(a) + y * cos(a)
func (w *Waypoint) cw(a int) {
	for i := 0; i < a/90; i++ {
		w.x, w.y = w.y, -w.x
	}
}

type ShipWaypoint struct {
	ship     Ship
	waypoint Waypoint
}

func (sw *ShipWaypoint) N(n int) {
	sw.waypoint.move(N, n)
}
func (sw *ShipWaypoint) S(n int) {
	sw.waypoint.move(S, n)
}
func (sw *ShipWaypoint) E(n int) {
	sw.waypoint.move(E, n)
}
func (sw *ShipWaypoint) W(n int) {
	sw.waypoint.move(W, n)
}
func (sw *ShipWaypoint) L(n int) {
	sw.waypoint.ccw(n)
}
func (sw *ShipWaypoint) R(n int) {
	sw.waypoint.cw(n)
}
func (sw *ShipWaypoint) F(n int) {
	sw.ship.E(sw.waypoint.x * n)
	sw.ship.N(sw.waypoint.y * n)
}

func execute(s Mover, instructions [][2]int) {
	for _, instruction := range instructions {
		n := instruction[1]
		switch instruction[0] {
		case N:
			s.N(n)
		case S:
			s.S(n)
		case W:
			s.W(n)
		case E:
			s.E(n)
		case L:
			s.L(n)
		case R:
			s.R(n)
		case F:
			s.F(n)
		}
	}
}

func parseInput(filepath string) (instructions [][2]int) {
	fp, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// Parse input
	scanner := bufio.NewScanner(fp)
	var action, value int
	for scanner.Scan() {
		line := scanner.Text()
		switch line[0] {
		case 'N':
			action = N
		case 'S':
			action = S
		case 'W':
			action = W
		case 'E':
			action = E
		case 'L':
			action = L
		case 'R':
			action = R
		case 'F':
			action = F
		}
		value, err = strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, [2]int{action, value})
	}
	return
}

func main() {
	var instructions [][2]int
	instructions = parseInput("input")

	// Part 1
	ship := Ship{0, 0, E}
	execute(&ship, instructions)
	fmt.Println(ship.taxiCabDist())

	// Part 2
	sw := ShipWaypoint{Ship{0, 0, E}, Waypoint{x: 10, y: 1}}
	execute(&sw, instructions)
	fmt.Println(sw.ship.taxiCabDist())
}
