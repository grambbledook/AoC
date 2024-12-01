package util

import (
	"github.com/grambbledook/adventofcode2023/2023/lang"
)

type Number interface {
	int | int64 | float64
}

func Max[N Number](a, b N) N {
	if a > b {
		return a
	}
	return b
}

func Sum[N Number](nums []N) N {
	return lang.Fold(nums, 0, func(a, b N) N { return a + b })
}

func Product[N Number](nums []N) N {
	return lang.Fold(nums, 1, func(a, b N) N { return a * b })
}

func Min[N Number](ch chan N) N {
	min := <-ch
	for n := range ch {
		if n < min {
			min = n
		}
	}
	return min
}

func Abs[N Number](num N) N {
	if num < 0 {
		return -num
	}
	return num
}
