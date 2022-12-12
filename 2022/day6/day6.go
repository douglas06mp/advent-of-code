package day6

import (
	"aoc2022/utils"
)

func First() int {
	input := utils.ParseInput("./day6/input.txt")
	max := 4

	return subStringWithoutRepeatingCharacters(input[0], max)
}

func Second() int {
	input := utils.ParseInput("./day6/input.txt")
	max := 14

	return subStringWithoutRepeatingCharacters(input[0], max)
}

func subStringWithoutRepeatingCharacters(str string, max int) int {
	l := 0
	seen := map[string]int{}

	for r := l; r < len(str); r++ {
		char := string(str[r])

		lastPosition, isExist := seen[char]
		if isExist && lastPosition >= l {
			l = lastPosition + 1
		}

		seen[char] = r
		if r-l+1 == max {
			return (r + 1)
		}
	}

	return -1
}
