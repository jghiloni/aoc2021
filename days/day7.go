package days

import (
	"fmt"
	"math"
	"sort"
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
	sort.Sort(sort.Reverse(sort.IntSlice(positionSet)))

	for position := 0; position <= positionSet[0]; position++ {
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
	// Attempt 1: Instead of multiplying by difference, multiply by arithmetic
	// subseries total. Remenber that
	//  k
	//  Σ  n = (n² + n)/2
	// n=1
	// EDIT: WRONG. Answer 97038219 is too high.

	// Attempt 2: I think I got lucky in Part 1 by only checking positions where
	// crabs already were. The example here shows a winning position that does
	// not have a crab already there. So, loop over all values between 0 and
	// max(positions) to see which wins.
	// Answer 97038163 is correct. Refactored Part 1 to use this method.

	crabPositions := slurpListAsInts(input7)
	positionSet, positionCounts := positionsAndCounts(crabPositions)

	fuelCount := math.MaxInt

	sort.Sort(sort.Reverse(sort.IntSlice(positionSet)))

	for position := 0; position <= positionSet[0]; position++ {
		tmpFuelCount := 0
		for k := range positionCounts {
			tmpFuelCount += positionCounts[k] * subsequenceTotal(intAbs(position-k))
		}

		if tmpFuelCount < fuelCount {
			fuelCount = tmpFuelCount
		}
	}

	return fmt.Sprint(fuelCount)
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

func subsequenceTotal(i int) int {
	return ((i * i) + i) / 2
}
