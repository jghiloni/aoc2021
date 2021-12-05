package days

import (
	"fmt"
	"math"
	"sort"
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
	bottomRight := getBoardSize(segments)

	board := make(map[point]int, bottomRight.x*bottomRight.y)

	for _, segment := range segments {
		if (segment.pointA.x != segment.pointB.x) && (segment.pointA.y != segment.pointB.y) {
			continue
		}

		if segment.pointA.x == segment.pointB.x {
			if segment.pointA.y > segment.pointB.y {
				segment.pointA, segment.pointB = segment.pointB, segment.pointA
			}

			for y := segment.pointA.y; y <= segment.pointB.y; y++ {
				board[point{segment.pointA.x, y}]++
			}
		}

		if segment.pointA.y == segment.pointB.y {
			if segment.pointA.x > segment.pointB.x {
				segment.pointA, segment.pointB = segment.pointB, segment.pointA
			}

			for x := segment.pointA.x; x <= segment.pointB.x; x++ {
				board[point{x, segment.pointA.y}]++
			}
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
	bottomRight := getBoardSize(segments)

	board := make(map[point]int, bottomRight.x*bottomRight.y)

	for _, segment := range segments {
		points := validateSegment(segment)
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

func getBoardSize(segments []lineSegment) point {
	points := make([]point, 0, len(segments))

	for _, seg := range segments {
		points = append(points, seg.pointA, seg.pointB)
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].x > points[j].x
	})

	maxX := points[0].x

	sort.Slice(points, func(i, j int) bool {
		return points[i].y > points[j].y
	})

	maxY := points[0].y

	return point{maxX, maxY}
}

func validateSegment(segment lineSegment) []point {
	deltaX, deltaY := segment.pointB.x-segment.pointA.x, segment.pointB.y-segment.pointA.y

	points := []point{}

	if deltaX == 0 {
		incr := deltaY / int(math.Abs(float64(deltaY)))
		for y := 0; y != deltaY+incr; y += incr {
			points = append(points, point{segment.pointA.x, segment.pointA.y + y})
		}

		return points
	}

	if deltaY == 0 {
		incr := deltaX / int(math.Abs(float64(deltaX)))
		for x := 0; x != deltaX+incr; x += incr {
			points = append(points, point{segment.pointA.x + x, segment.pointA.y})
		}

		return points
	}

	if math.Abs(float64(deltaX)) == math.Abs(float64(deltaY)) {
		incr := int(math.Abs(float64(deltaX)))
		xSign := deltaX / int(math.Abs(float64(deltaX)))
		ySign := deltaY / int(math.Abs(float64(deltaY)))

		for i := 0; i <= incr; i++ {
			points = append(points, point{segment.pointA.x + (i * xSign), segment.pointA.y + (i * ySign)})
		}

		return points
	}

	return nil
}
