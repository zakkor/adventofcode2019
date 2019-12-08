package main

import (
	"strconv"
	"strings"
)

func Day1Part1() int {
	modules := day1ProcessedInput()

	var totalFuel int
	for _, mod := range modules {
		// Add module weight
		totalFuel += mod/3 - 2
	}
	return totalFuel
}

func Day1Part2() int {
	modules := day1ProcessedInput()

	var totalFuel int
	for _, mod := range modules {
		// Add module weight
		fuel := mod/3 - 2
		totalFuel += fuel

		// Add fuel weight
		for {
			fuel = fuel/3 - 2
			if fuel <= 0 {
				break
			}
			totalFuel += fuel
		}
	}
	return totalFuel
}

func day1ProcessedInput() []int {
	input := readInput("./day1.txt")
	var nums []int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}
