package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type day4 struct{}

type item struct {
	value  int
	marked bool
}

type board [5][5]item

type boards []board

func (d *day4) Part1() string {
	boards, calledNumbers := getBoards()

	winningBoard, lastInput := findFirstWinner(boards, calledNumbers)

	return fmt.Sprintf("winning board:\n%v\n\nscore: %d", winningBoard, score(winningBoard, lastInput))
}

func (d *day4) Part2() string {
	boards, calledNumbers := getBoards()

	winningBoard, lastInput := findLastWinner(boards, calledNumbers)

	return fmt.Sprintf("winning board:\n%v\n\nscore: %d", winningBoard, score(winningBoard, lastInput))
}

func getBoards() (boards, []string) {
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

	return boards, calledNumbers
}

func findFirstWinner(b boards, calledNumbers []string) (board, int) {

	for i := range calledNumbers {
		calledNumber, err := strconv.Atoi(calledNumbers[i])
		if err != nil {
			log.Fatal(err)
		}

		for j := range b {
			if mark(&b[j], calledNumber) {
				return b[j], calledNumber
			}
		}
	}

	return board([5][5]item{}), -1
}

func findLastWinner(b boards, calledNumbers []string) (board, int) {
	completedBoards := make([]int, 0, len(b))
	lastNumber := -1

	for j := range calledNumbers {
		calledNumber, err := strconv.Atoi(calledNumbers[j])
		if err != nil {
			log.Fatal(err)
		}

		for i := range b {
			if indexOf(completedBoards, i) != -1 {
				continue
			}

			if mark(&b[i], calledNumber) {
				lastNumber = calledNumber
				completedBoards = append(completedBoards, i)
			}
		}
	}

	return b[completedBoards[len(completedBoards)-1]], lastNumber
}

func indexOf(b []int, val int) int {
	for i := range b {
		if b[i] == val {
			return i
		}
	}

	return -1
}

func mark(b *board, num int) bool {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if b[row][col].value == num {
				b[row][col].marked = true
				if check(*b) {
					return true
				}
			}
		}
	}

	return false
}

func check(b board) bool {
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

func score(b board, multiplier int) int {
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

func (b board) String() string {
	return fmt.Sprintf("%s %s %s %s %s\n%s %s %s %s %s\n%s %s %s %s %s\n%s %s %s %s %s\n%s %s %s %s %s",
		colorize(b[0][0]), colorize(b[0][1]), colorize(b[0][2]), colorize(b[0][3]), colorize(b[0][4]),
		colorize(b[1][0]), colorize(b[1][1]), colorize(b[1][2]), colorize(b[1][3]), colorize(b[1][4]),
		colorize(b[2][0]), colorize(b[2][1]), colorize(b[2][2]), colorize(b[2][3]), colorize(b[2][4]),
		colorize(b[3][0]), colorize(b[3][1]), colorize(b[3][2]), colorize(b[3][3]), colorize(b[3][4]),
		colorize(b[4][0]), colorize(b[4][1]), colorize(b[4][2]), colorize(b[4][3]), colorize(b[4][4]),
	)
}

func colorize(i item) string {
	str := fmt.Sprintf("%02d", i.value)
	if i.marked {
		str = color.GreenString(str)
	}

	return str
}
