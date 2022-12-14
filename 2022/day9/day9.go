package day9

import (
	"aoc2022/utils"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

const (
	LEGNTH = 1000
	MID    = LEGNTH / 2
)

func First() int {
	input := utils.ParseInput("./day9/input.txt")

	visited := makeBoard(LEGNTH)

	count := 0
	H := Position{x: MID, y: MID}
	T := Position{x: MID, y: MID}

	for _, action := range input {
		direction := strings.Split(action, " ")[0]
		step, _ := strconv.Atoi(strings.Split(action, " ")[1])

		for i := 0; i < step; i++ {
			moveH(&H, direction, 1)
			moveT(&T, &H)

			if !visited[T.x][T.y] {
				count++
				visited[T.x][T.y] = true
			}
		}
	}

	return count
}

func Second() int {
	input := utils.ParseInput("./day9/input.txt")

	visited := makeBoard(LEGNTH)
	count := 0

	knots := make([]Position, 10)
	for i := range knots {
		knots[i] = Position{x: MID, y: MID}
	}

	for _, action := range input {
		direction := strings.Split(action, " ")[0]
		step, _ := strconv.Atoi(strings.Split(action, " ")[1])

		for i := 0; i < step; i++ {
			moveH(&knots[0], direction, 1)
			for i := 1; i < len(knots); i++ {
				moveT(&knots[i], &knots[i-1])
			}

			end := knots[9]
			if !visited[end.x][end.y] {
				count++
				visited[end.x][end.y] = true
			}
		}
	}

	return count
}

func moveH(H *Position, dir string, step int) {
	switch dir {
	// UP
	case "U":
		(*H).y++
	// RIGHT
	case "R":
		(*H).x++
	// DOWN
	case "D":
		(*H).y--
	// LEFT
	case "L":
		(*H).x--
	}
}

func moveT(T *Position, H *Position) {
	xDiff := int(math.Abs(float64(H.x) - float64(T.x)))
	yDiff := int(math.Abs(float64(H.y) - float64(T.y)))

	switch true {
	// RIGHT, LEFT
	case xDiff == 2 && yDiff == 0:
		if H.x > T.x {
			(*T).x++
		} else {
			(*T).x--
		}
	// UP, DOWN
	case xDiff == 0 && yDiff == 2:
		if H.y > T.y {
			(*T).y++
		} else {
			(*T).y--
		}
	// twelve CORNER
	case xDiff == 2 && yDiff == 1, xDiff == 1 && yDiff == 2, xDiff == 2 && yDiff == 2:
		if H.x > T.x && H.y > T.y {
			(*T).x++
			(*T).y++
		}
		if H.x > T.x && H.y < T.y {
			(*T).x++
			(*T).y--
		}
		if H.x < T.x && H.y < T.y {
			(*T).x--
			(*T).y--
		}
		if H.x < T.x && H.y > T.y {
			(*T).x--
			(*T).y++
		}
	}
}

func makeBoard(length int) [][]bool {
	board := make([][]bool, length)
	for i := range board {
		board[i] = make([]bool, length)
	}

	return board
}
