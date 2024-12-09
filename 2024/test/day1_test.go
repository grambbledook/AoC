package test

import (
	"github.com/grambbledook/adventofcode2023/2024/day1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day1_Task1_TestData(t *testing.T) {
	result := day1.Task1("../day1/test1.data")
	assert.Equal(t, 11, result)
}

func Test_Day1_Task1_InputData(t *testing.T) {
	result := day1.Task1("../day1/input.data")
	assert.Equal(t, 1110981, result)
}

func Test_Day1_Task2_TestData(t *testing.T) {
	result := day1.Task2("../day1/test1.data")
	assert.Equal(t, 31, result)
}

func Test_Day1_Task2_InputData(t *testing.T) {
	result := day1.Task2("../day1/input.data")
	assert.Equal(t, 24869388, result)
}
