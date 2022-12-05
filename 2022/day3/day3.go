package day3

import (
	"aoc2022/utils"
	"strings"
)

var priorities = map[int]int{}

func initPriorities() {
	for i := 0; i < 26; i++ {
		priorities['a'+i] = i + 1
		priorities['A'+i] = i + 27
	}
}

func First() int {
	initPriorities()
	input := utils.ParseInput("./day3/input.txt")

	sum := 0
	for _, str := range input {
		middle := len(str) / 2
		first := str[:middle]
		second := str[middle:]

		charMap := map[string]int{}
		for _, char := range strings.Split(first, "") {
			if _, isExist := charMap[char]; !isExist {
				charMap[char] = 1
			}
		}

		for _, char := range second {
			if _, isExist := charMap[string(char)]; isExist {
				sum += priorities[int(char)]
				break
			}
		}
	}

	return sum
}

func Second() int {
	initPriorities()
	input := utils.ParseInput("./day3/input.txt")

	sum := 0
	for i := 0; i < len(input); i += 3 {
		set2, set3 := toSet(input[i+1]), toSet(input[i+2])

		for _, char := range input[i] {
			if set2[string(char)] && set3[string(char)] {
				sum += priorities[int(char)]
				break
			}
		}
	}

	return sum
}

func toSet(s string) map[string]bool {
	set := map[string]bool{}
	for _, char := range strings.Split(s, "") {
		set[char] = true
	}

	return set
}
