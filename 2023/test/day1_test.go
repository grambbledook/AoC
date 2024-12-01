package test

import (
	"github.com/grambbledook/adventofcode2023/2023/day1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day1_1_TestData(t *testing.T) {
	result := day1.Compute(day1.Task1, "../day1/test1")
	assert.Equal(t, 142, result)
}

func Test_Day1_1_TaskData(t *testing.T) {
	result := day1.Compute(day1.Task1, "../day1/input")
	assert.Equal(t, 54968, result)
}

func Test_Day1_2_TestData(t *testing.T) {
	result := day1.Compute(day1.Task2, "../day1/test2")
	assert.Equal(t, 281, result)
}

func Test_Day1_2_TaskData(t *testing.T) {
	result := day1.Compute(day1.Task2, "../day1/input")
	assert.Equal(t, 54094, result)
}
