package test

import (
	"github.com/grambbledook/adventofcode2023/2024/day5"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day5_Task1_TestData(t *testing.T) {
	result := day5.Task1("../day5/test1")
	assert.Equal(t, 143, result)
}

func Test_Day5_Task1_InputData(t *testing.T) {
	result := day5.Task1("../day5/input")
	assert.Equal(t, 4689, result)
}

func Test_Day5_Task2_TestData(t *testing.T) {
	result := day5.Task2("../day5/test1")
	assert.Equal(t, 123, result)
}

func Test_Day5_Task2_InputData(t *testing.T) {
	result := day5.Task2("../day5/input")
	assert.Equal(t, 6336, result)
}
