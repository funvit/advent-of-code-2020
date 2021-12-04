package day1

import (
	"testing"

	"aoc2020"
)

func TestPart1(t *testing.T) {
	type args struct {
		lines       []string
		expectedSum int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"sample",
			args{
				lines:       aoc2020.MustReadLinesFromFile("../input/sample1.txt"),
				expectedSum: 2020,
			},
			514579,
			false,
		},
		{
			"input",
			args{
				lines:       aoc2020.MustReadLinesFromFile("../input/day1.txt"),
				expectedSum: 2020,
			},
			357504,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part1(tt.args.lines, tt.args.expectedSum)
			if (err != nil) != tt.wantErr {
				t.Errorf("Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Part1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		lines       []string
		expectedSum int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"sample",
			args{
				lines:       aoc2020.MustReadLinesFromFile("../input/sample1.txt"),
				expectedSum: 2020,
			},
			241861950,
			false,
		},
		{
			"input",
			args{
				lines:       aoc2020.MustReadLinesFromFile("../input/day1.txt"),
				expectedSum: 2020,
			},
			12747392,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part2(tt.args.lines, tt.args.expectedSum)
			if (err != nil) != tt.wantErr {
				t.Errorf("Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Part2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
