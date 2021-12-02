package days

import (
	"fmt"
	"log"
	"strings"
)

type day2 struct{}

func (d *day2) Part1() string {
	lines := slurpListAsLines(input2)

	var (
		direction string
		val       int
	)

	h := 0
	v := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		fmt.Sscanf(line, "%s %d", &direction, &val)

		switch strings.ToLower(direction) {
		case "forward":
			h += val
		case "down":
			v += val
		case "up":
			v -= val
		default:
			log.Fatalf("unknown command %s", direction)
		}
	}

	return fmt.Sprintf("%d", h*v)
}

func (d *day2) Part2() string {
	lines := slurpListAsLines(input2)

	var (
		direction string
		val       int
	)

	h := 0
	v := 0
	aim := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		fmt.Sscanf(line, "%s %d", &direction, &val)

		switch strings.ToLower(direction) {
		case "forward":
			h += val
			v += val * aim
		case "down":
			aim += val
		case "up":
			aim -= val
		default:
			log.Fatalf("unknown command %s", direction)
		}
	}

	return fmt.Sprintf("%d", h*v)
}
