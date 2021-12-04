package day3

type Slope struct {
	Right, Down int
}

// Part1 .
//
// Returns count of trees "#" by slope right and down.
func Part1(lines []string, slope Slope) int {
	var row, col, i int

	maxCol := len(lines[0])

	for row < len(lines) {
		if lines[row][col] == '#' {
			i++
		}

		row += slope.Down
		col += slope.Right

		col = col % maxCol
	}

	return i
}

func Part2(lines []string, slopes []Slope) int {

	var r int

	for _, slope := range slopes {
		v := Part1(lines, slope)
		if r == 0 {
			r = v
			continue
		}
		r *= v
	}

	return r
}
