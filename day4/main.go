package main

import (
	"AoC2022/utils"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type Range struct {
	Start int
	End   int
}

func (r1 *Range) Contains(r2 Range) bool {
	return r2.Start >= r1.Start && r2.End <= r1.End
}

func (r1 *Range) Overlaps(r2 Range) bool {
	return r1.Start <= r2.End && r1.End >= r2.Start
}

type Assignment struct {
	Elf1 Range
	Elf2 Range
}

func (a *Assignment) HasFullyContainedRange() bool {
	return a.Elf1.Contains(a.Elf2) || a.Elf2.Contains(a.Elf1)
}

func (a *Assignment) HasOverlappingRange() bool {
	return a.Elf1.Overlaps(a.Elf2)
}
func GetAssignments() []Assignment {
	lines := strings.Split(input, "\n")
	assignments := make([]Assignment, len(lines))

	for idx, l := range lines {
		elfRangeStrings := strings.Split(l, ",")
		elf1Strings := strings.Split(elfRangeStrings[0], "-")
		elf2Strings := strings.Split(elfRangeStrings[1], "-")

		assignments[idx] = Assignment{
			Range{utils.MustInt(elf1Strings[0]), utils.MustInt(elf1Strings[1])},
			Range{utils.MustInt(elf2Strings[0]), utils.MustInt(elf2Strings[1])},
		}
	}
	return assignments
}

func main() {
	part1 := 0
	part2 := 0
	for _, a := range GetAssignments() {
		if a.HasFullyContainedRange() {
			part1++
		}
		if a.HasOverlappingRange() {
			part2++
		}
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
