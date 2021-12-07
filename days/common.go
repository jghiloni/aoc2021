package days

import (
	"log"
	"strconv"
	"strings"
)

func slurpListAsInts(input string) []int {
	list := strings.Fields(strings.ReplaceAll(input, ",", " "))

	ints := make([]int, len(list))

	for i := range list {
		x, e := strconv.Atoi(list[i])
		if e != nil {
			log.Fatal(e)
		}

		ints[i] = x
	}

	return ints
}

func slurpListAsBinaryInts(input string) []int {
	list := strings.Fields(input)

	ints := make([]int, len(list))

	for i := range list {
		x, e := strconv.ParseInt(list[i], 2, 32)
		if e != nil {
			log.Fatal(e)
		}

		ints[i] = int(x)
	}

	return ints
}

func slurpListAsLines(input string) []string {
	return strings.Split(input, "\n")
}

func filterInts(list []int, filterFn func(int) bool) []int {
	filtered := make([]int, 0, len(list))
	for i := range list {
		if filterFn(list[i]) {
			filtered = append(filtered, list[i])
		}
	}

	return filtered
}
