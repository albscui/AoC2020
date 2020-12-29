package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	FLOOR    = '.'
	OCCUPIED = '#'
	EMPTY    = 'L'
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

func (b1 Board) equal(b2 Board) bool {
	for r := range b1 {
		for c := range b1[r] {
			if b1[r][c] != b2[r][c] {
				return false
			}
		}
	}
	return true
}

func (b Board) adjacentCells(r, c int) (result []rune) {
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

func (b Board) searchInDirection(r, c int, d [2]int) (ans rune) {
	dr, dc := d[0], d[1]
	_r, _c := r+dr, c+dc
	for b.validCoord(_r, _c) {
		if b[_r][_c] != FLOOR {
			return b[_r][_c]
		}
		_r, _c = _r+dr, _c+dc
	}
	return FLOOR
}

func (b Board) firstSeatsInAllDirections(r, c int) (output []rune) {
	directions := [8][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, d := range directions {
		output = append(output, b.searchInDirection(r, c, d))
	}
	return
}

func (b Board) validCoord(r, c int) bool {
	return 0 <= r && r < len(b) && 0 <= c && c < len(b[0])
}

func (b Board) countOccupiedAjacent(r, c int) (ans int) {
	for _, adj := range b.adjacentCells(r, c) {
		if adj == OCCUPIED {
			ans++
		}
	}
	return
}

func (b Board) countOccupiedAllDirections(r, c int) (ans int) {
	for _, s := range b.firstSeatsInAllDirections(r, c) {
		if s == OCCUPIED {
			ans++
		}
	}
	return
}

/*
If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
*/
func (b Board) nextCellState(cell rune, r int, c int) rune {
	if cell == EMPTY && b.countOccupiedAjacent(r, c) == 0 {
		return OCCUPIED
	}

	if cell == OCCUPIED && b.countOccupiedAjacent(r, c) >= 4 {
		return EMPTY
	}
	return cell
}

func (b Board) nextCellStateV2(cell rune, r, c int) rune {
	if cell == EMPTY && b.countOccupiedAllDirections(r, c) == 0 {
		return OCCUPIED
	}

	if cell == OCCUPIED && b.countOccupiedAllDirections(r, c) >= 5 {
		return EMPTY
	}
	return cell
}

func (b Board) countOccupied() (ans int) {
	for r := range b {
		for c := range b[r] {
			if b[r][c] == OCCUPIED {
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

func (b Board) nextV2() Board {
	b2 := Board{}
	for r, row := range b {
		newRow := []rune{}
		for c, cell := range row {
			newRow = append(newRow, b.nextCellStateV2(cell, r, c))
		}
		b2 = append(b2, newRow)
	}
	return b2
}

func (b Board) copy() (b2 Board) {
	b2 = make(Board, len(b))
	for i := range b {
		b2[i] = make([]rune, len(b[i]))
		copy(b2[i], b[i])
	}
	return
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

	// Part 1
	grid1 := grid.copy()
	nextGrid := grid1.next()
	for !grid1.equal(nextGrid) {
		grid1, nextGrid = nextGrid, nextGrid.next()
	}
	fmt.Println(grid1.countOccupied())

	// Part 2
	grid2 := grid.copy()
	nextGrid = grid2.nextV2()
	for !grid2.equal(nextGrid) {
		grid2, nextGrid = nextGrid, nextGrid.nextV2()
	}
	fmt.Println(grid2.countOccupied())
}
