package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
	//scores := map[string]int{
	//	"A X": 1 + 3, // rock + rock
	//	"A Y": 2 + 6, // rock + paper
	//	"A Z": 3 + 0, // rock + scissors
	//	"B X": 1 + 0, // paper + rock
	//	"B Y": 2 + 3, // paper + paper
	//	"B Z": 3 + 6, // paper + scissors
	//	"C X": 1 + 6, // scissors + rock
	//	"C Y": 2 + 0, // scissors + paper
	//	"C Z": 3 + 3, // scissors + scissors
	//}
	scores := map[string]int{
		"A X": 3 + 0, // scissors + lose
		"A Y": 1 + 3, // rock + draw
		"A Z": 2 + 6, // rock + win
		"B X": 1 + 0, // paper + lose
		"B Y": 2 + 3, // paper + draw
		"B Z": 3 + 6, // paper + win
		"C X": 2 + 0, // scissors + lose
		"C Y": 3 + 3, // scissors + draw
		"C Z": 1 + 6, // scissors + win
	}
	score := 0
	for _, line := range strings.Split(input, "\n") {
		score += scores[line]
	}
	fmt.Println(score)
}
