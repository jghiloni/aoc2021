package days

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
)

type day3 struct{}

func (d *day3) Part1() string {
	readings := slurpListAsBinaryInts(input3)
	if len(readings) < 1 {
		log.Fatal("no input")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(readings)))

	msbPosition := int(math.Floor(math.Log2(float64(readings[0]))))
	γ := 0
	ε := 0
	majority := len(readings) / 2
	for bit := msbPosition; bit >= 0; bit-- {
		countOnes := 0
		bitPositionVal := 1 << bit
		for _, reading := range readings {
			if reading&bitPositionVal == bitPositionVal {
				countOnes++
			}
		}

		if countOnes > majority {
			γ = γ | bitPositionVal
		} else {
			ε = ε | bitPositionVal
		}
	}

	return fmt.Sprintf("γ = %d [%s], ε = %d [%s], W = %d", γ, strconv.FormatInt(int64(γ), 2), ε, strconv.FormatInt(int64(ε), 2), γ*ε)
}

func (d *day3) Part2() string {
	readings := slurpListAsBinaryInts(input3)
	if len(readings) < 1 {
		log.Fatal("no input")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(readings)))

	msbPosition := int(math.Floor(math.Log2(float64(readings[0]))))

	filteredO2 := make([]int, len(readings))
	filteredCO2 := make([]int, len(readings))

	copy(filteredO2, readings)
	copy(filteredCO2, readings)

	filterOnOne := func(bitVal int) func(x int) bool {
		return func(x int) bool {
			return x&bitVal == bitVal
		}
	}

	filterOnZero := func(bitVal int) func(x int) bool {
		return func(x int) bool {
			return x&bitVal == 0
		}
	}

	for bit := msbPosition; len(filteredO2) > 1 && bit >= 0; bit-- {
		majority := len(filteredO2) / 2
		countOnes := 0
		bitPositionVal := 1 << bit
		for _, reading := range filteredO2 {
			if reading&bitPositionVal == bitPositionVal {
				countOnes++
			}
		}

		if countOnes >= majority {
			filteredO2 = filterInts(filteredO2, filterOnOne(bitPositionVal))
		} else {
			filteredO2 = filterInts(filteredO2, filterOnZero(bitPositionVal))
		}
	}

	for bit := msbPosition; len(filteredCO2) > 1 && bit >= 0; bit-- {
		majority := len(filteredCO2) / 2
		countZeroes := 0
		bitPositionVal := 1 << bit
		for _, reading := range filteredCO2 {
			if reading&bitPositionVal == 0 {
				countZeroes++
			}
		}

		if countZeroes <= majority {
			filteredCO2 = filterInts(filteredCO2, filterOnZero(bitPositionVal))
		} else {
			filteredCO2 = filterInts(filteredCO2, filterOnOne(bitPositionVal))
		}
	}

	return fmt.Sprintf("O2 Generator Rating = %d, CO2 Scrubber Rating = %d, Life Support Rating = %d", filteredO2[0], filteredCO2[0], filteredO2[0]*filteredCO2[0])
}
