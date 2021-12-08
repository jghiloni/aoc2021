package days

import (
	"fmt"
	"strings"
)

type day8 struct{}

func (d *day8) Part1() string {
	// Part 1's algorithm is pretty simple, and is hard not to derive directly
	// from the text. For each line, throw away the input values and look for
	// the values that correspond to a 1, 4, 7, or 8. Because the example
	// explicitly calls out that each of these digits uses a unique number of
	// digits (1 uses 2, 4 uses 4, 7 uses 3, and 8 uses 7), simply check each
	// output value for one of those string lengths and increment if found. The
	// answer 321 is correct.

	lines := slurpListAsLines(input8)

	importantNumberCount := 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Split(line, "|")
		for _, outputDigit := range strings.Fields(parts[1]) {
			switch len(outputDigit) {
			case 2: // 1
				fallthrough
			case 3: // 7
				fallthrough
			case 4: // 4
				fallthrough
			case 7: // 8
				importantNumberCount++
			}
		}
	}

	return fmt.Sprint(importantNumberCount)
}

func (d *day8) Part2() string {
	return ""
}
