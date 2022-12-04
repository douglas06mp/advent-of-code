package day2

import (
	"aoc2022/utils"
	"strings"
)

var scoreMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

const (
	Win  = 6
	Draw = 3
	Loss = 0
)

func First() int {
	input := utils.ParseInput("./day2/input.txt")

	charMap := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	score := 0
	for _, round := range input {
		opponent := strings.Split(round, " ")[0]
		me := strings.Split(round, " ")[1]

		me = charMap[me]
		score += scoreMap[me]

		switch me {
		case "A":
			switch opponent {
			case "A":
				score += Draw
			case "B":
				score += Loss
			case "C":
				score += Win
			}
		case "B":
			switch opponent {
			case "A":
				score += Win
			case "B":
				score += Draw
			case "C":
				score += Loss
			}

		case "C":
			switch opponent {
			case "A":
				score += Loss
			case "B":
				score += Win
			case "C":
				score += Draw
			}
		}
	}

	return score
}

func Second() int {
	input := utils.ParseInput("./day2/input.txt")

	resultMap := map[string]int{
		"X": Loss,
		"Y": Draw,
		"Z": Win,
	}

	score := 0
	for _, round := range input {
		opponent := strings.Split(round, " ")[0]
		result := resultMap[strings.Split(round, " ")[1]]

		score += result

		switch result {
		case Win:
			switch opponent {
			case "A":
				score += scoreMap["B"]
			case "B":
				score += scoreMap["C"]
			case "C":
				score += scoreMap["A"]
			}
		case Loss:
			switch opponent {
			case "A":
				score += scoreMap["C"]
			case "B":
				score += scoreMap["A"]
			case "C":
				score += scoreMap["B"]
			}

		case Draw:
			switch opponent {
			case "A":
				score += scoreMap["A"]
			case "B":
				score += scoreMap["B"]
			case "C":
				score += scoreMap["C"]
			}
		}
	}

	return score
}
