package day1

import (
	"aoc2022/utils"
	"sort"
	"strconv"
)

func First() int {
	input := utils.ParseInput("./day1/input.txt")
	nums := group(input)

	max := 0
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}

func Second() int {
	input := utils.ParseInput("./day1/input.txt")
	nums := group(input)

	max := []int{0, 0, 0}
	for _, n := range nums {
		if n > max[2] {
			max[2] = n
			sort.Slice(max, func(a, b int) bool {
				return max[a] > max[b]
			})
		}
	}

	sum := 0
	for _, m := range max {
		sum += m
	}
	return sum
}

func group(input []string) []int {
	var nums []int
	temp := 0

	for _, v := range input {
		if len(v) == 0 {
			nums = append(nums, temp)
			temp = 0
		} else {
			value, _ := strconv.Atoi(v)
			temp += value
		}
	}

	nums = append(nums, temp)
	return nums
}
