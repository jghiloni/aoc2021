package days

import (
	"fmt"
	"sort"
	"strings"
)

type day10 struct{}

const (
	openRunes  = "([{<"
	closeRunes = ")]}>"
)

var (
	errorPrizes = []int{3, 57, 1197, 25137}
)

func (d *day10) Part1() string {
	lines := slurpListAsLines(input10)
	totalScore := 0
	for _, line := range lines {
		totalScore += getLineValue(line)
	}

	return fmt.Sprint(totalScore)
}

func (d *day10) Part2() string {
	autocompleteScores := []int{}

	lines := slurpListAsLines(input10)
	for _, line := range lines {
		if score := getAutocompleteLineScore(line); score > 0 {
			autocompleteScores = append(autocompleteScores, score)
		}
	}

	sort.Ints(autocompleteScores)

	return fmt.Sprint(autocompleteScores[len(autocompleteScores)/2])
}

func getLineValue(line string) int {
	stack := make([]rune, 0, len(line)/2)

	for _, r := range line {
		if strings.ContainsRune(openRunes, r) {
			stack = append(stack, r)
			continue
		}

		if idx := strings.IndexRune(closeRunes, r); idx != -1 {
			if len(stack) > 0 && stack[len(stack)-1] == rune(openRunes[idx]) {
				stack = stack[:len(stack)-1] // pop the last item off since it matched
				continue
			}

			// if we're here it means that it's corrupt, not just incomplete
			return errorPrizes[idx]
		}
	}

	return 0
}

func getAutocompleteLineScore(line string) int {
	stack := make([]rune, 0, len(line)/2)

	for _, r := range line {
		if strings.ContainsRune(openRunes, r) {
			stack = append(stack, r)
			continue
		}

		if idx := strings.IndexRune(closeRunes, r); idx != -1 {
			if len(stack) > 0 && stack[len(stack)-1] == rune(openRunes[idx]) {
				stack = stack[:len(stack)-1] // pop the last item off since it matched
				continue
			}

			// if we're here it means that it's corrupt, not incomplete
			return 0
		}
	}

	// there will be remaining characters in the stack here, so get each
	// character, and find its closing counterpart
	total := 0
	for i := len(stack) - 1; i >= 0; i-- {
		popped := stack[i]
		idx := strings.IndexRune(openRunes, popped)

		total *= 5
		total += idx + 1
	}

	return total
}
