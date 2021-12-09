package days

import (
	"fmt"
	"strconv"
)

type day9 struct{}

func (d *day9) Part1() string {
	matrix := [][]byte{}
	rows := slurpListAsLines(input9)

	for _, row := range rows {
		matrix = append(matrix, []byte(row))
	}

	total := 0
	for r := range rows {
		for c := range rows[r] {
			if isLowPoint(matrix, r, c) {
				lp, _ := strconv.Atoi(string(matrix[r][c]))
				total += lp + 1
			}
		}
	}
	return fmt.Sprint(total)
}

func (d *day9) Part2() string {
	return ""
}

func isLowPoint(matrix [][]byte, r, c int) bool {
	var (
		left  byte = 'A'
		right byte = 'A'
		up    byte = 'A'
		down  byte = 'A'
	)

	if c > 0 {
		left = matrix[r][c-1]
	}

	if c < len(matrix[r])-1 {
		right = matrix[r][c+1]
	}

	if r > 0 {
		up = matrix[r-1][c]
	}

	if r < len(matrix)-1 {
		down = matrix[r+1][c]
	}

	b := matrix[r][c]
	if b < left && b < right && b < up && b < down {
		return true
	}

	return false
}
