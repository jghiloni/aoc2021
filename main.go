package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jghiloni/aoc2021/days"
)

func main() {
	var (
		day      uint
		exercise uint
	)

	flag.UintVar(&day, "day", 1, "Which day's problem to execute (defaults to 1)")
	flag.UintVar(&exercise, "exercise", 1, "Which exercise to run (1 or 2, defaults to 1)")
	flag.Parse()

	dayIntf, ok := days.DayMap[day]
	if !ok {
		log.Fatalf("day %d has not been implemented", day)
	}

	if exercise < 1 || exercise > 2 {
		log.Fatalf("you must choose exercise 1 or 2, you chose %d", exercise)
	}

	output := ""
	if exercise == 1 {
		output = dayIntf.Part1()
	} else {
		output = dayIntf.Part2()
	}

	fmt.Println(output)
}
