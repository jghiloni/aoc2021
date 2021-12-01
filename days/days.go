package days

var DayMap = make(map[string]Day, 24)

type Day interface {
	Part1() string
	Part2() string
}

func init() {
	DayMap["1"] = &day1{}
}
