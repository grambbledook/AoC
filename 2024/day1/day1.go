package day1

import (
	"github.com/grambbledook/adventofcode2023/assert"
	"github.com/grambbledook/adventofcode2023/util"
	"sort"
	"strconv"
	"strings"
)

func Task1(file string) int {
	lines := util.ReadLines(file)

	var result int

	var left []int
	var right []int

	for _, line := range lines {
		nums := strings.Split(line, "\t")

		l, err := strconv.Atoi(nums[0])
		assert.NoError(err)

		r, err := strconv.Atoi(nums[1])
		assert.NoError(err)

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := range left {
		result += util.Abs(left[i] - right[i])
	}

	return result
}

func Task2(file string) int {
	lines := util.ReadLines(file)

	var result int

	var left = make([]int, 0)
	var right = make(map[int]int)

	for _, line := range lines {
		nums := strings.Split(line, "\t")

		l, err := strconv.Atoi(nums[0])
		assert.NoError(err)

		r, err := strconv.Atoi(nums[1])
		assert.NoError(err)

		left = append(left, l)
		right[r] += 1
	}

	for i, n := range left {
		result += left[i] * right[n]
	}

	return result
}
