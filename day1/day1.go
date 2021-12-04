package day1

import (
	"fmt"
	"strconv"
)

// Part1 .
//
// Returns -1 if expected sum not found.
func Part1(lines []string, expectedSum int64) (int64, error) {

	// pre-filter values
	filteredInts := make([]int64, 0, len(lines))
	for i, v := range lines {
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("parse line %d: %w", i+1, err)
		}

		if n > expectedSum {
			// too high
			continue
		}
		filteredInts = append(filteredInts, n)
	}

	for xi, x := range filteredInts {
		// insane. non-optimal
		for yi, y := range filteredInts {
			if yi == xi {
				continue
			}
			if x+y == expectedSum {
				return x * y, nil
			}
		}
	}

	return -1, nil
}

// Part2 .
//
// Returns -1 if expected sum not found.
func Part2(lines []string, expectedSum int64) (int64, error) {

	// pre-filter values
	filteredInts := make([]int64, 0, len(lines))
	for i, v := range lines {
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("parse line %d: %w", i+1, err)
		}

		if n > expectedSum {
			// too high
			continue
		}
		filteredInts = append(filteredInts, n)
	}

	for xi, x := range filteredInts {
		// insane. non-optimal
		for yi, y := range filteredInts {
			if yi == xi {
				continue
			}
			// omg!
			for zi, z := range filteredInts {
				if zi == yi || zi == xi {
					continue
				}
				if x+y+z == expectedSum {
					return x * y * z, nil
				}
			}
		}
	}

	return -1, nil
}
