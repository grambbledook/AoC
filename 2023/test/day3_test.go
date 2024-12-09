package test

import (
	"github.com/grambbledook/adventofcode2023/2023/day3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day3_1_TestData(t *testing.T) {
	result := day3.Compute(day3.Task1, "../day3/test.data")
	assert.Equal(t, 4361, result)
}

func Test_Day3_1_TaskData(t *testing.T) {
	result := day3.Compute(day3.Task1, "../day3/input.data")
	assert.Equal(t, 559667, result)
}

func Test_Day3_2_TestData(t *testing.T) {
	result := day3.Compute(day3.Task2, "../day3/test.data")
	assert.Equal(t, 467835, result)
}

func Test_Day3_2_TaskData(t *testing.T) {
	result := day3.Compute(day3.Task2, "../day3/input.data")
	assert.Equal(t, 86841457, result)
}
