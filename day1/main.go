package main

import (
	"AoC2022/utils"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {

	elfs := strings.Split(input, "\n\n")

	var allElfCalories []int
	for _, elf := range elfs {
		calorieStrings := strings.Split(elf, "\n")
		elfCalories := 0
		for _, calorieString := range calorieStrings {
			calorie, err := strconv.Atoi(calorieString)
			utils.CheckError(err)
			elfCalories += calorie
		}
		allElfCalories = append(allElfCalories, elfCalories)
	}

	sort.Ints(allElfCalories)
	fmt.Println(allElfCalories[len(allElfCalories)-1])

	part2 := 0
	for _, c := range allElfCalories[len(allElfCalories)-3:] {
		part2 += c
	}
	fmt.Println(part2)
}
