package main

import (
	"flag"
	"fmt"
	"os"

	"aoc2020"
	"aoc2020/day1"
	"aoc2020/day2"
	"aoc2020/day3"
	"aoc2020/day4"
)

func main() {
	fmt.Println("Advent of code 2020")

	var (
		day, part uint
	)
	flag.UintVar(&day, "day", 1, "day of advent")
	flag.UintVar(&part, "part", 1, "part of day")
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage of %s:\n",
			os.Args[0],
		)
		flag.PrintDefaults()

		fmt.Fprintf(
			flag.CommandLine.Output(),
			"\nExample: \n\n\t%s -day 1 -part 1 ./input.1-1.txt\n\n",
			os.Args[0],
		)
	}

	flag.Parse()

	//
	// process
	//
	inFile := flag.Arg(0)
	if inFile == "" {
		fmt.Println("ERROR: specify input data file as last arg")
		os.Exit(1)
	}

	f, err := os.Open(inFile)
	if err != nil {
		fmt.Println("ERROR: file open:", err)
		os.Exit(1)
	}
	defer f.Close()

	lines, err := aoc2020.ReadLines(f)
	if err != nil {
		fmt.Println("ERROR: read lines:", err)
		os.Exit(1)
	}

	fmt.Println("Day:", day, "Part:", part)
	fmt.Println("Lines:", len(lines))

	switch day {
	case 1:
		callPart(
			part-1,
			func() {
				r, err := day1.Part1(lines, 2020)
				if err != nil {
					printErrAndExit(err.Error())
				}
				printAnswerAndExit(r)
			},
			func() {
				r, err := day1.Part2(lines, 2020)
				if err != nil {
					printErrAndExit(err.Error())
				}
				printAnswerAndExit(r)
			},
		)

	case 2:
		callPart(
			part-1,
			func() {
				r := day2.AssertPolicy(lines, func(s string) day2.IPolicy {
					return day2.NewParsePolicy1(s)
				})
				printAnswerAndExit(r)
			},
			func() {
				r := day2.AssertPolicy(lines, func(s string) day2.IPolicy {
					return day2.NewParsePolicy2(s)
				})
				printAnswerAndExit(r)
			},
		)
	case 3:
		callPart(
			part-1,
			func() {
				r := day3.Part1(lines, day3.Slope{Right: 3, Down: 1})
				printAnswerAndExit(r)
			},
			func() {

				r := day3.Part2(lines, []day3.Slope{
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
				})
				printAnswerAndExit(r)
			},
		)

	case 4:
		callPart(
			part-1,
			func() {
				r := day4.CountValidPassports(lines, day4.ValidatePassport)
				printAnswerAndExit(r)
			},
			func() {
				r := day4.CountValidPassports(lines, day4.ValidatePassportStrict)
				printAnswerAndExit(r)
			},
		)

	default:
		fmt.Println("unknown day")
	}
}

func callPart(index uint, fn ...func()) {
	if index >= uint(len(fn)) {
		fmt.Println("unknown part")
		return
	}
	fn[index]()
}

func printErrAndExit(args ...interface{}) {
	fmt.Println("ERROR", args)
	os.Exit(1)
}

func printAnswerAndExit(val interface{}) {
	fmt.Println("Answer:", val)
	os.Exit(0)
}
