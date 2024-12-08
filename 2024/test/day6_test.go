package test

import (
	"github.com/grambbledook/adventofcode2023/2024/day6"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day6_Task1_TestData(t *testing.T) {
	result := day6.Task1("../day6/test1")
	assert.Equal(t, 41, result)
}

func Test_Day6_Task1_InputData(t *testing.T) {
	result := day6.Task1("../day6/input")
	assert.Equal(t, 5199, result)
}

func Test_Day6_Task2_TestData(t *testing.T) {
	result := day6.Task2("../day6/test1")
	assert.Equal(t, 6, result)
}

func Test_Day6_Task2_InputData(t *testing.T) {
	result := day6.Task2("../day6/input")
	assert.Equal(t, 1915, result)
}
