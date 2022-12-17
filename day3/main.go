package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type Rucksack struct {
	compartment1 string
	compartment2 string
}

type Group struct {
	elfs []Rucksack
}

func (r *Rucksack) Items() string {
	return r.compartment1 + r.compartment2
}

func (r *Rucksack) FindDuplicate() rune {
	for _, c := range r.compartment1 {
		if strings.ContainsRune(r.compartment2, c) {
			return c
		}
	}
	panic(fmt.Sprintf("no duplicate in %v", r))
}

func GetRucksacks() []Rucksack {
	lines := strings.Split(input, "\n")

	rucksacks := make([]Rucksack, 0, len(lines))
	for _, line := range lines {
		rucksacks = append(rucksacks, Rucksack{
			line[:len(line)/2],
			line[len(line)/2:],
		})
	}
	return rucksacks
}

func GetGroups() []Group {
	rucksacks := GetRucksacks()
	groups := make([]Group, 0, len(rucksacks)/3)
	for i := 0; i < len(rucksacks); i += 3 {
		groups = append(groups, Group{rucksacks[i : i+3]})
	}
	return groups
}

func (g *Group) GetCommonItem() rune {
	for _, c := range g.elfs[0].Items() {
		matches := 0
		for _, elf := range g.elfs[1:] {
			if strings.ContainsRune(elf.Items(), c) {
				matches++
			}
		}
		if matches == len(g.elfs)-1 {
			return c
		}
	}
	panic(fmt.Sprintf("no common item in %v", g))
}
func itemToScore(c rune) int {
	if int(c) > 90 {
		// lower case
		return int(c) - 96
	} else {
		// upper case
		return int(c) - 38
	}
}

func main() {
	score := 0
	for _, r := range GetRucksacks() {
		score += itemToScore(r.FindDuplicate())
	}

	fmt.Printf("Part 1: %d\n", score)

	part2 := 0
	for _, g := range GetGroups() {
		part2 += itemToScore(g.GetCommonItem())
	}

	fmt.Printf("Part 2: %d\n", part2)
}
