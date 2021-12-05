package days

import (
	"fmt"
)

type day5 struct{}

type point struct {
	x int
	y int
}

type lineSegment struct {
	pointA point
	pointB point
}

func (d *day5) Part1() string {
	segments := readInput()
	board := make(map[point]int)

	for _, segment := range segments {
		points := validateSegment(segment, false)
		for _, p := range points {
			board[p]++
		}
	}

	count := 0
	for _, v := range board {
		if v >= 2 {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}

func (d *day5) Part2() string {
	segments := readInput()
	board := make(map[point]int)

	for _, segment := range segments {
		points := validateSegment(segment, true)
		for _, p := range points {
			board[p]++
		}
	}

	count := 0
	for _, v := range board {
		if v >= 2 {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}

func readInput() []lineSegment {
	lines := slurpListAsLines(input5)

	segments := make([]lineSegment, 0, len(lines))

	for _, line := range lines {
		var (
			a point
			b point
		)

		fmt.Sscanf(line, "%d,%d -> %d,%d", &a.x, &a.y, &b.x, &b.y)
		segments = append(segments, lineSegment{pointA: a, pointB: b})
	}

	return segments
}

func validateSegment(segment lineSegment, diagonalAllowed bool) []point {
	deltaX, deltaY := segment.pointB.x-segment.pointA.x, segment.pointB.y-segment.pointA.y

	points := []point{}

	if deltaX == 0 {
		incr := deltaY / intAbs(deltaY)
		for y := 0; y != deltaY+incr; y += incr {
			points = append(points, point{segment.pointA.x, segment.pointA.y + y})
		}

		return points
	}

	if deltaY == 0 {
		incr := deltaX / intAbs(deltaX)
		for x := 0; x != deltaX+incr; x += incr {
			points = append(points, point{segment.pointA.x + x, segment.pointA.y})
		}

		return points
	}

	if diagonalAllowed && (intAbs(deltaX) == intAbs(deltaY)) {
		incr := intAbs(deltaX)
		xSign := deltaX / intAbs(deltaX)
		ySign := deltaY / intAbs(deltaY)

		for i := 0; i <= incr; i++ {
			points = append(points, point{segment.pointA.x + (i * xSign), segment.pointA.y + (i * ySign)})
		}

		return points
	}

	return nil
}

func intAbs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}
