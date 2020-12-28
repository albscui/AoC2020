package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	floor    = '.'
	occupied = '#'
	empty    = 'L'
)

type cell struct {
	row   int
	col   int
	state int
}

type Board [][]rune

func (b Board) myPrint() {
	for _, row := range b {
		fmt.Println(string(row))
	}
}

func (b1 Board) isSame(b2 Board) bool {
	for r := range b1 {
		for c := range b1[r] {
			if b1[r][c] != b2[r][c] {
				return false
			}
		}
	}
	return true
}

func (b Board) adjacentCells(cell rune, r, c int) (result []rune) {
	// up, down, left, right, ul, ur, dl, dr
	directions := [8][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, d := range directions {
		dr, dc := d[0], d[1]
		_r, _c := r+dr, c+dc
		if b.validCoord(_r, _c) {
			result = append(result, b[_r][_c])
		}
	}
	return
}

func (b Board) validCoord(r, c int) bool {
	return 0 <= r && r < len(b) && 0 <= c && c < len(b[0])
}

func (b Board) noOccupiedAdjacent(cell rune, r, c int) bool {
	for _, adj := range b.adjacentCells(cell, r, c) {
		if adj == occupied {
			return false
		}
	}
	return true
}

func (b Board) fourOrMoreOccupied(cell rune, r, c int) bool {
	numOccupied := 0
	for _, adj := range b.adjacentCells(cell, r, c) {
		if adj == occupied {
			numOccupied++
		}
	}
	return numOccupied >= 4
}

/*
If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
*/
func (b Board) nextCellState(cell rune, r int, c int) rune {
	if cell == empty && b.noOccupiedAdjacent(cell, r, c) {
		return occupied
	} else if cell == occupied && b.fourOrMoreOccupied(cell, r, c) {
		return empty
	}
	return cell
}

func (b Board) countOccupied() (ans int) {
	for r := range b {
		for c := range b[r] {
			if b[r][c] == occupied {
				ans++
			}
		}
	}
	return
}

func (b Board) next() Board {
	b2 := Board{}
	for r, row := range b {
		newRow := []rune{}
		for c, cell := range row {
			newRow = append(newRow, b.nextCellState(cell, r, c))
		}
		b2 = append(b2, newRow)
	}
	return b2
}

func main() {
	fp, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	grid := Board{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	nextGrid := grid.next()
	for !grid.isSame(nextGrid) {
		grid, nextGrid = nextGrid, nextGrid.next()
	}
	fmt.Println(grid.countOccupied())
}
