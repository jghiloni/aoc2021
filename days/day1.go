package days

import (
	"fmt"
	"log"
	"math"
)

type day1 struct{}

func (d *day1) Part1() string {
	distances := slurpListAsInts(input1)
	if len(distances) < 1 {
		log.Fatal("missing input")
	}

	prevDistance := distances[0]
	numIncreases := 0
	for _, currentDistance := range distances {

		if currentDistance > prevDistance {
			numIncreases++
		}

		prevDistance = currentDistance
	}

	return fmt.Sprintf("%d", numIncreases)
}

func (d *day1) Part2() string {
	distances := slurpListAsInts(input1)
	if len(distances) < 1 {
		log.Fatal("missing input")
	}

	numIncreases := 0
	prevTotal := math.MaxInt
	for i := range distances {
		currentTotal := get3dTotal(distances, i)

		log.Printf("%d %d %d", numIncreases, prevTotal, currentTotal)
		if currentTotal == -1 {
			break
		}

		if currentTotal > prevTotal {
			numIncreases++
		}

		prevTotal = currentTotal
	}

	return fmt.Sprintf("%d", numIncreases)
}

func get3dTotal(distances []int, startIdx int) int {
	if startIdx > len(distances)-3 {
		return -1
	}

	return distances[startIdx] + distances[startIdx+1] + distances[startIdx+2]
}
