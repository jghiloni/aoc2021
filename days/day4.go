package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type day4 struct{}

type item struct {
	value  int
	marked bool
}

type board [5][5]item

type boards []board

func (d *day4) Part1() string {
	lines := slurpListAsLines(input4)

	calledNumbers := strings.Split(lines[0], ",")

	lines = lines[2:]

	boards := make(boards, 0, len(lines)/6)
	for boardCounter := 0; boardCounter < len(lines); boardCounter += 6 {
		board := board([5][5]item{})

		for rowCounter := 0; rowCounter < 5; rowCounter++ {
			vals := strings.Fields(lines[boardCounter+rowCounter])
			row := [5]item{}
			for i := range vals {
				val, err := strconv.Atoi(vals[i])
				if err != nil {
					log.Fatal(err)
				}
				row[i] = item{value: val}
			}
			board[rowCounter] = row
		}

		boards = append(boards, board)
	}

	winningBoard, lastInput := boards.findFirstWinner(calledNumbers)

	return fmt.Sprintf("winning board: %v\n\nscore: %d", winningBoard, winningBoard.score(lastInput))
}

func (d *day4) Part2() string {
	return ""
}

func (b boards) findFirstWinner(calledNumbers []string) (board, int) {
	var winningBoard board

	for i := range calledNumbers {
		calledNumber, err := strconv.Atoi(calledNumbers[i])
		if err != nil {
			log.Fatal(err)
		}

		for _, board := range b {
			if board.mark(calledNumber) {
				return winningBoard, calledNumber
			}
		}
	}

	return board([5][5]item{}), -1
}

func (b *board) mark(num int) bool {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if b[row][col].value == num {
				b[row][col].marked = true
				if b.check() {
					return true
				}
			}
		}
	}

	return false
}

func (b board) check() bool {
	for i := 0; i < 5; i++ {
		// check row i
		if b[i][0].marked &&
			b[i][1].marked &&
			b[i][2].marked &&
			b[i][3].marked &&
			b[i][4].marked {
			return true
		}

		// check column i
		if b[0][i].marked &&
			b[1][i].marked &&
			b[2][i].marked &&
			b[3][i].marked &&
			b[4][i].marked {
			return true
		}
	}

	return false
}

func (b board) score(multiplier int) int {
	sum := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !b[row][col].marked {
				sum += b[row][col].value
			}
		}
	}

	return sum * multiplier
}
