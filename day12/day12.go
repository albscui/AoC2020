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

func getAction(a byte) int {
	switch a {
	case 'N':
		return N
	case 'S':
		return S
	case 'W':
		return W
	case 'E':
		return E
	case 'L':
		return L
	case 'R':
		return R
	case 'F':
		return F
	}
	return F
}

type Ship struct {
	x, y      int
	direction int
}

func (s *Ship) move(a, v int) {
	switch a {
	case N:
		s.y += v
	case S:
		s.y -= v
	case E:
		s.x += v
	case W:
		s.x -= v
	case L, R:
		s.rotate(a, v)
	case F:
		s.move(s.direction, v)
	}
}

func (s *Ship) rotate(d, a int) {
	// 0, 90, 180, 270, 360
	var directions [4]int
	switch d {
	case L:
		directions = [4]int{E, N, W, S}
	case R:
		directions = [4]int{E, S, W, N}
	}
	for i, d := range directions {
		if s.direction == d {
			_idx := int(math.Mod(float64(i+a/90), 4))
			s.direction = directions[_idx]
			break
		}
	}
}

func (s Ship) taxiCabDist() float64 {
	return math.Abs(float64(s.x)) + math.Abs(float64(s.y))
}

type Waypoint struct {
	x, y int
}

func (w *Waypoint) move(a, v int) {
	switch a {
	case N:
		w.y += v
	case S:
		w.y -= v
	case W:
		w.x -= v
	case E:
		w.x += v
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

func (sw *ShipWaypoint) rotateWaypoint(d, a int) {
	switch d {
	case L:
		sw.waypoint.ccw(a)
	case R:
		sw.waypoint.cw(a)
	}
}

func (sw *ShipWaypoint) moveShipToWaypoint(n int) {
	for i := 0; i < n; i++ {
		sw.ship.x += sw.waypoint.x
		sw.ship.y += sw.waypoint.y
	}
}

func (sw *ShipWaypoint) move(action, value int) {
	switch action {
	case N:
		sw.waypoint.move(action, value)
	case S:
		sw.waypoint.move(action, value)
	case E:
		sw.waypoint.move(action, value)
	case W:
		sw.waypoint.move(action, value)
	case L, R:
		sw.rotateWaypoint(action, value)
	case F:
		sw.moveShipToWaypoint(value)
	}

}

func execute(s *Ship, instructions [][2]int) {
	for _, instruction := range instructions {
		s.move(instruction[0], instruction[1])
	}
}

func execute2(s *ShipWaypoint, instructions [][2]int) {
	for _, instruction := range instructions {
		s.move(instruction[0], instruction[1])
	}
}

func main() {
	fp, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// Parse input
	instructions := [][2]int{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		action := getAction(line[0])
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, [2]int{action, value})
	}

	// Part 1
	ship := Ship{0, 0, E}
	execute(&ship, instructions)
	fmt.Println(ship.taxiCabDist())

	// Part 2
	sw := ShipWaypoint{Ship{0, 0, E}, Waypoint{x: 10, y: 1}}
	execute2(&sw, instructions)
	fmt.Println(sw.ship.taxiCabDist())
}
