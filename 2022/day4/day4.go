package day4

import (
	"aoc2022/utils"
	"strconv"
	"strings"
)

func First() int {
	input := utils.ParseInput("./day4/input.txt")

	sum := 0
	for _, floor := range input {
		first, second := split(floor, ",")
		firstStart, firstEnd := splitAndInt(first, "-")
		secondStart, secondEnd := splitAndInt(second, "-")

		if ((firstStart <= secondStart) && (firstEnd >= secondEnd)) ||
			((firstStart >= secondStart) && (firstEnd <= secondEnd)) {
			sum++
		}
	}
	return sum
}

func Second() int {
	input := utils.ParseInput("./day4/input.txt")

	sum := 0
	for _, floor := range input {
		first, second := split(floor, ",")
		firstStart, firstEnd := splitAndInt(first, "-")
		secondStart, secondEnd := splitAndInt(second, "-")

		if ((firstStart <= secondStart) && (firstEnd >= secondStart)) ||
			((firstStart >= secondStart) && (firstStart <= secondEnd)) {
			sum++
		}
	}
	return sum
}

func split(s string, sep string) (string, string) {
	split := strings.Split(s, sep)
	return split[0], split[1]
}

func splitAndInt(s string, sep string) (int, int) {
	split := strings.Split(s, sep)
	first, _ := strconv.Atoi(split[0])
	second, _ := strconv.Atoi(split[1])

	return first, second
}
