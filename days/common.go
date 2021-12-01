package days

import (
	"log"
	"strconv"
	"strings"
)

func slurpListAsInts(input string) []int {
	list := strings.Fields(input)

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
