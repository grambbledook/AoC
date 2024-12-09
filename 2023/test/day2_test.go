package test

import (
	"github.com/grambbledook/adventofcode2023/2023/day2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day2_1_TestData(t *testing.T) {
	result := day2.Compute(day2.Task1, "../day2/test.data")
	assert.Equal(t, 8, result)
}

func Test_Day2_1_TaskData(t *testing.T) {
	result := day2.Compute(day2.Task1, "../day2/input.data")
	assert.Equal(t, 2239, result)
}

func Test_Day2_2_TestData(t *testing.T) {
	result := day2.Compute(day2.Task2, "../day2/test.data")
	assert.Equal(t, 2286, result)
}

func Test_Day2_2_TaskData(t *testing.T) {
	result := day2.Compute(day2.Task2, "../day2/input.data")
	assert.Equal(t, 83435, result)
}
