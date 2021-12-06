package days

import (
	"fmt"
	"strings"
)

type day6 struct{}

func (d *day6) Part1() string {
	return countFishies(80)
}

func (d *day6) Part2() string {
	return countFishies(256)
}

func countFishies(numDays int) string {
	fishies := parseInput()

	for i := 0; i < numDays; i++ {
		processDay(fishies)
	}

	return fmt.Sprint(totalSum(fishies))
}

func parseInput() map[int]int64 {
	fishies := map[int]int64{}
	for _, remainingGestationPeriod := range slurpListAsInts(strings.ReplaceAll(input6, ",", " ")) {
		fishies[remainingGestationPeriod]++
	}
	return fishies
}

func processDay(fishies map[int]int64) {
	birthingFishies, hasZeroes := fishies[0]
	for i := 1; i <= 8; i++ {
		fishies[i-1] = fishies[i]
	}

	if hasZeroes {
		fishies[6] += birthingFishies
		fishies[8] = birthingFishies
	}
}

func totalSum(fishies map[int]int64) int64 {
	sum := int64(0)
	for _, v := range fishies {
		sum += v
	}

	return sum
}
