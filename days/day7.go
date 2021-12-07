package days

import "fmt"

type day7 struct{}

func (d *day7) Part1() string {
	// initial theory based on example input: normalizing to the mode provides
	// the optimal fuel consumption as it provides the most zeroes

	crabPositions := slurpListAsInts(input7)
	normalPosition := mode(crabPositions)

	fuelTotal := 0
	for _, pos := range crabPositions {
		fuelTotal += intAbs(pos - normalPosition)
	}

	return fmt.Sprintf("mode = %d, fuel total = %d", normalPosition, fuelTotal)
}

func (d *day7) Part2() string {
	return ""
}

func mode(list []int) int {
	set := map[int]int{}

	for _, i := range list {
		set[i]++
	}

	max := 0
	for k := range set {
		if set[k] > set[max] {
			max = k
		}
	}

	return max
}
