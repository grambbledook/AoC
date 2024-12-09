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
		out[i] = Int(str)
	}

	return out
}

func Uints(input string) []uint64 {
	in := strings.Fields(input)
	out := make([]uint64, len(in))

	for i, str := range in {
		out[i] = Uint(str)
	}

	return out
}

func Int(str string) int {
	num, err := strconv.Atoi(str)
	assert.NoError(err)

	return num
}

func Uint(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	assert.NoError(err)

	return num
}
