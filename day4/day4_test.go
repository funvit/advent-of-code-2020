package day4

import (
	"testing"

	"aoc2020"
)

func TestCountValidPassports(t *testing.T) {
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
			args{lines: aoc2020.MustReadLinesFromFile("../input/sample4.txt")},
			2,
		},
		{
			"input",
			args{lines: aoc2020.MustReadLinesFromFile("../input/day4.txt")},
			233,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountValidPassports(tt.args.lines, ValidatePassport); got != tt.want {
				t.Errorf("CountValidPassports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountValidPassportsStrict(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"sample invalid",
			args{lines: aoc2020.MustReadLinesFromFile("../input/sample4.invalid.txt")},
			0,
		},
		{
			"sample valid",
			args{lines: aoc2020.MustReadLinesFromFile("../input/sample4.valid.txt")},
			4,
		},
		{
			"input",
			args{lines: aoc2020.MustReadLinesFromFile("../input/day4.txt")},
			111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountValidPassports(tt.args.lines, ValidatePassportStrict); got != tt.want {
				t.Errorf("CountValidPassports() = %v, want %v", got, tt.want)
			}
		})
	}
}
