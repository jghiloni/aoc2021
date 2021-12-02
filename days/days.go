package days

var DayMap = make(map[uint]Day, 24)

type Day interface {
	Part1() string
	Part2() string
}

func init() {
	DayMap[1] = &day1{}
	DayMap[2] = &day2{}
}
