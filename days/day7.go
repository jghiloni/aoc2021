package days

import (
	"fmt"
	"math"
)

type day7 struct{}

func (d *day7) Part1() string {
	// initial theory based on example input: normalizing to the mode provides
	// the optimal fuel consumption as it provides the most zeroes. O(n)
	// EDIT: WRONG. Answer 425971 is too high

	// Attempt 2: Convert list to map[num]countOfNum. O(n + m^2), n = max(list).
	// Answer 345035 is correct

	crabPositions := slurpListAsInts(input7)
	positionSet, positionCounts := positionsAndCounts(crabPositions)

	fuelCount := math.MaxInt
	for _, position := range positionSet {
		tmpFuelCount := 0
		for k := range positionCounts {
			tmpFuelCount += positionCounts[k] * intAbs(position-k)
		}

		if tmpFuelCount < fuelCount {
			fuelCount = tmpFuelCount
		}
	}

	return fmt.Sprint(fuelCount)
}

func (d *day7) Part2() string {
	return ""
}

func positionsAndCounts(list []int) ([]int, map[int]int) {
	set := map[int]int{}

	for _, i := range list {
		set[i]++
	}

	positions := make([]int, 0, len(set))
	for k := range set {
		positions = append(positions, k)
	}

	return positions, set
}
