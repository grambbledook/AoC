package day2

import (
	"github.com/grambbledook/adventofcode2023/util"
)

func Task1(file string) int {
	lines := util.ReadLines(file)

	var result int

	for _, line := range lines {
		nums := util.Ints(line)

		isValid := Test(nums)

		if isValid {
			result++
		}
	}

	return result
}

func Task2(file string) int {
	lines := util.ReadLines(file)

	var result int

	for _, line := range lines {
		nums := util.Ints(line)

		isValid := Test(nums)

		if isValid {
			result++
			continue
		}

		for i := range nums {
			fixed := make([]int, 0, len(nums)-1)
			fixed = append(fixed, nums[:i]...)
			fixed = append(fixed, nums[i+1:]...)

			isValid = Test(fixed)
			if isValid {
				result++
				break
			}
		}
	}

	return result
}

func Test(nums []int) bool {

	isAscending := true
	isDescending := true

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		if util.Abs(diff) < 1 || util.Abs(diff) > 3 {
			return false
		}

		if diff < 0 {
			isAscending = false
		}

		if diff > 0 {
			isDescending = false
		}
	}
	return isAscending || isDescending
}
