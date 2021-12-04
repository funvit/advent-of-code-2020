package day3

import (
	"testing"

	"aoc2020"
)

func TestPart1(t *testing.T) {
	type args struct {
		lines []string
		slope Slope
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"sample",
			args{
				lines: aoc2020.MustReadLinesFromFile("../input/sample3.txt"),
				slope: Slope{
					Right: 3,
					Down:  1,
				},
			},
			7,
		},
		{
			"input",
			args{
				lines: aoc2020.MustReadLinesFromFile("../input/day3.txt"),
				slope: Slope{
					Right: 3,
					Down:  1,
				},
			},
			218,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.lines, tt.args.slope); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	slopes := []Slope{
		{
			Right: 1,
			Down:  1,
		},
		{
			Right: 3,
			Down:  1,
		},
		{
			Right: 5,
			Down:  1,
		},
		{
			Right: 7,
			Down:  1,
		},
		{
			Right: 1,
			Down:  2,
		},
	}

	type args struct {
		lines  []string
		slopes []Slope
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"sample",
			args{
				lines:  aoc2020.MustReadLinesFromFile("../input/sample3.txt"),
				slopes: slopes,
			},
			336,
		},
		{
			"input",
			args{
				lines:  aoc2020.MustReadLinesFromFile("../input/day3.txt"),
				slopes: slopes,
			},
			3847183340,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.lines, tt.args.slopes); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
