package main

import (
	"AoC2022/utils"
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"
)

//go:embed example
var input string

type Monkey struct {
	Items          []int
	Operation      func(int) int
	TestModulo     int
	IfTrueMonkey   int
	IfFalseMonkey  int
	InspectedItems int
}

func Add(n1, n2 int) int {
	return n1 + n2
}

func Mul(n1, n2 int) int {
	return n1 * n2
}

func ParseOp(s string) func(int) int {
	var operands []string
	var operator func(int, int) int
	if strings.Contains(s, "+") {
		operands = strings.Split(s, "+")
		operator = Add
	} else {
		operands = strings.Split(s, "*")
		operator = Mul
	}

	operands[0] = strings.TrimSpace(operands[0])
	operands[1] = strings.TrimSpace(operands[1])

	if operands[0] == "old" && operands[1] == "old" {
		return func(old int) int {
			return operator(old, old)
		}
	}
	if operands[0] == "old" {
		return func(old int) int {
			return operator(old, utils.MustInt(operands[1]))
		}
	}
	return func(old int) int {
		return operator(utils.MustInt(operands[0]), utils.MustInt(operands[1]))
	}
}
func GetMonkeys() []*Monkey {
	monkeys := make([]*Monkey, 0)
	lines := strings.Split(input, "\n")

	for lineIdx := 0; lineIdx < len(lines); {
		items := strings.Split(lines[lineIdx+1][len("  Starting items: "):], ", ")
		intItems := utils.StrSliceToInt(items)

		op := ParseOp(lines[lineIdx+2][len("  Operation: new = "):])
		test := utils.MustInt(lines[lineIdx+3][len("  Test: divisible by "):])
		iftrue := utils.MustInt(lines[lineIdx+4][len("    If true: throw to monkey "):])
		iffalse := utils.MustInt(lines[lineIdx+5][len("   If false: throw to monkey "):])

		m := &Monkey{
			intItems,
			op,
			test,
			iftrue,
			iffalse,
			0,
		}
		monkeys = append(monkeys, m)
		lineIdx += 7
	}
	return monkeys
}

func Round(ms []*Monkey) {
	for _, m := range ms {
		for idx, i := range m.Items {
			m.InspectedItems += 1
			m.Items[idx] = int(math.Floor(float64(m.Operation(i)) / 3))

			var nextMonkey int
			if m.Items[idx]%m.TestModulo == 0 {
				nextMonkey = m.IfTrueMonkey
			} else {
				nextMonkey = m.IfFalseMonkey
			}
			ms[nextMonkey].Items = append(ms[nextMonkey].Items, m.Items[idx])
		}
		m.Items = m.Items[:0]

	}
}
func main() {
	ms := GetMonkeys()
	for round := 0; round < 20; round++ {
		Round(ms)
	}
	inspections := make([]int, 0)

	for _, m := range ms {
		inspections = append(inspections, m.InspectedItems)
		//fmt.Printf("%v\n", m.InspectedItems)
	}
	sort.Ints(inspections)
	part1 := inspections[len(inspections)-2] * inspections[len(inspections)-1]
	fmt.Printf("Part 1: %d", part1)

}
