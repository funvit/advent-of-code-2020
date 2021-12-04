package day5

import (
	"fmt"
)

const (
	maxRow = 127
	maxCol = 7
)

const (
	front = 'F'
	back  = 'B'
	left  = 'L'
	right = 'R'
)

func Part1(lines []string) int {
	var ids []int
	var max int

	for _, l := range lines {
		row, col := Pos(l)
		id := ID(row, col)
		ids = append(ids, id)

		if id > max {
			max = id
		}
	}

	return max
}

func Part2(lines []string) int {

	// row => col => taken
	seats := make([][]bool, maxRow+1)
	for r := range seats {
		seats[r] = make([]bool, maxCol+1)
		for c := range seats[r] {
			seats[r][c] = false
		}
	}

	for _, l := range lines {
		r, c := Pos(l)
		seats[r][c] = true
	}

	// debug: print seats
	//for r := range seats {
	//	var b strings.Builder
	//	for c := range seats[r] {
	//		if seats[r][c] {
	//			b.WriteRune('X')
	//			continue
	//		}
	//		b.WriteRune('.')
	//	}
	//	fmt.Println(b.String())
	//}

	pos := func(i int) (r, c int) {
		c = i % (maxCol + 1)
		r = i / (maxCol + 1)
		return r, c
	}

	for i := maxCol + 1; i < (maxRow+1)*(maxCol+1); i++ {
		r, c := pos(i)

		if r == 0 {
			continue
		}

		prevR, prevC := pos(i - 1)
		nextR, nextC := pos(i + 1)

		if !seats[r][c] && seats[prevR][prevC] && seats[nextR][nextC] {
			fmt.Println("My seat: row:", r, "col:", c)
			return ID(r, c)
		}
	}

	return 0
}

func ID(row, col int) int {
	return row*8 + col
}

func Pos(s string) (row, col int) {
	rowStart := 0
	rowEnd := maxRow
	colStart := 0
	colEnd := maxCol

	for _, c := range s {
		switch c {
		case front:
			rowStart, rowEnd = lowerHalf(rowStart, rowEnd)
		case back:
			rowStart, rowEnd = upperHalf(rowStart, rowEnd)
		case left:
			colStart, colEnd = lowerHalf(colStart, colEnd)
		case right:
			colStart, colEnd = upperHalf(colStart, colEnd)
		}
	}

	return rowEnd, colEnd
}

func lowerHalf(start, end int) (start_, end_ int) {
	return start, (start + end) / 2
}

func upperHalf(start, end int) (start_, end_ int) {
	return (start + end) / 2, end
}
