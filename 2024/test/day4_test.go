package test

import (
	"github.com/grambbledook/adventofcode2023/2024/day4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day4_Task1_TestData(t *testing.T) {
	result := day4.Task1("../day4/test1")
	assert.Equal(t, 18, result)
}

func Test_Day4_Task1_InputData(t *testing.T) {
	result := day4.Task1("../day4/input")
	assert.Equal(t, 2297, result)
}

func Test_Day4_Task2_TestData(t *testing.T) {
	result := day4.Task2("../day4/test1")
	assert.Equal(t, 9, result)
}

func Test_Day4_Task2_TestData_Lean(t *testing.T) {
	result := day4.Task2("../day4/test.lean")
	assert.Equal(t, 9, result)
}

func Test_Day4_Task2_InputData(t *testing.T) {
	result := day4.Task2("../day4/input")
	assert.Equal(t, 1745, result)
}
