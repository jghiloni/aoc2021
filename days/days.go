package days

var DayMap = make(map[uint]Day, 24)

type Day interface {
	Part1() string
	Part2() string
}

func init() {
	DayMap[1] = &day1{}
	DayMap[2] = &day2{}
	DayMap[3] = &day3{}
	DayMap[4] = &day4{}
	DayMap[5] = &day5{}
	DayMap[6] = &day6{}
	DayMap[7] = &day7{}
	DayMap[8] = &day8{}
	DayMap[9] = &day9{}
	DayMap[10] = &day10{}
	DayMap[11] = &day11{}
	DayMap[12] = &day12{}
	DayMap[13] = &day13{}
	DayMap[14] = &day14{}
	DayMap[15] = &day15{}
	DayMap[16] = &day16{}
	DayMap[17] = &day17{}
	DayMap[18] = &day18{}
	DayMap[19] = &day19{}
	DayMap[20] = &day20{}
	DayMap[21] = &day21{}
	DayMap[22] = &day22{}
	DayMap[23] = &day23{}
	DayMap[24] = &day24{}
	DayMap[25] = &day25{}
}
