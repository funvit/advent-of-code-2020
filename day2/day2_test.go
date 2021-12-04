package day2

import (
	"testing"

	"aoc2020"
)

func TestPart1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"sample",
			args{lines: aoc2020.MustReadLinesFromFile("../input/sample2.txt")},
			2,
		},
		{
			"input",
			args{lines: aoc2020.MustReadLinesFromFile("../input/day2.txt")},
			467,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssertPolicy(tt.args.lines, func(s string) IPolicy {
				return NewParsePolicy1(s)
			}); got != tt.want {
				t.Errorf("AssertPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"sample",
			args{lines: aoc2020.MustReadLinesFromFile("../input/sample2.txt")},
			1,
		},
		{
			"input",
			args{lines: aoc2020.MustReadLinesFromFile("../input/day2.txt")},
			441,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssertPolicy(tt.args.lines, func(s string) IPolicy {
				return NewParsePolicy2(s)
			}); got != tt.want {
				t.Errorf("AssertPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}
