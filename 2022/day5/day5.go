package day5

import (
	"aoc2022/utils"
	"regexp"
	"strconv"
)

var crate1 = []string{"B", "W", "N"}
var crate2 = []string{"L", "Z", "S", "P", "T", "D", "M", "B"}
var crate3 = []string{"Q", "H", "Z", "W", "R"}
var crate4 = []string{"W", "D", "V", "J", "Z", "R"}
var crate5 = []string{"S", "H", "M", "B"}
var crate6 = []string{"L", "G", "N", "J", "H", "V", "P", "B"}
var crate7 = []string{"J", "Q", "Z", "F", "H", "D", "L", "S"}
var crate8 = []string{"W", "S", "F", "J", "G", "Q", "B"}
var crate9 = []string{"Z", "W", "M", "S", "C", "D", "J"}

func First() string {
	input := utils.ParseInput("./day5/input.txt")

	crates := map[string][]string{
		"1": crate1,
		"2": crate2,
		"3": crate3,
		"4": crate4,
		"5": crate5,
		"6": crate6,
		"7": crate7,
		"8": crate8,
		"9": crate9,
	}

	for _, command := range input {
		move, from, to := parseCommand(command)
		for i := 0; i < move; i++ {
			crates[to] = append(crates[to], crates[from][len(crates[from])-1])
			crates[from] = crates[from][:len(crates[from])-1]
		}
	}

	msg := ""
	for i := 1; i < 10; i++ {
		c := crates[strconv.Itoa(i)]
		msg += c[len(c)-1]
	}

	return msg
}

func Second() string {
	input := utils.ParseInput("./day5/input.txt")

	crates := map[string][]string{
		"1": crate1,
		"2": crate2,
		"3": crate3,
		"4": crate4,
		"5": crate5,
		"6": crate6,
		"7": crate7,
		"8": crate8,
		"9": crate9,
	}

	for _, command := range input {
		move, from, to := parseCommand(command)

		diff := len(crates[from]) - move
		crates[to] = append(crates[to], crates[from][diff:]...)
		crates[from] = crates[from][:diff]

	}

	msg := ""
	for i := 1; i < 10; i++ {
		c := crates[strconv.Itoa(i)]
		msg += c[len(c)-1]
	}

	return msg
}

func parseCommand(command string) (int, string, string) {
	r, _ := regexp.Compile(`^move (\d+) from (\d+) to (\d+)`)
	matches := r.FindStringSubmatch(command)

	move, _ := strconv.Atoi(matches[1])
	from := matches[2]
	to := matches[3]

	return move, from, to
}
