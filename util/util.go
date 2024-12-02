package util

import (
	"github.com/grambbledook/adventofcode2023/assert"
	"strconv"
	"strings"
)

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Ints(input string) []int {
	in := strings.Fields(input)
	out := make([]int, len(in))

	for i, str := range in {
		num, err := strconv.Atoi(str)
		assert.NoError(err)

		out[i] = num
	}

	return out
}
