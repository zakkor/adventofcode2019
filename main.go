package main

import (
	"flag"
	"fmt"
)

var days = map[string]func() int{
	"1-1": Day1Part1,
	"1-2": Day1Part2,
	"2-1": Day2Part1,
	"2-2": Day2Part2,
}

func main() {
	var day, part int
	flag.IntVar(&day, "day", 0, "")
	flag.IntVar(&part, "part", 0, "")
	flag.Parse()

	if day == 0 {
		fmt.Println("specify a day using -day")
		return
	}
	if part == 0 {
		fmt.Println("specify a part using -part")
		return
	}

	key := fmt.Sprintf("%v-%v", day, part)
	if compute, ok := days[key]; ok {
		fmt.Println(compute())
	} else {
		fmt.Println("no solution for key", key)
	}
}
