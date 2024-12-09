package test

import (
	"github.com/grambbledook/adventofcode2023/2024/day7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day7_Task1_TestData(t *testing.T) {
	result := day7.Task1("../day7/test1.data")
	assert.Equal(t, 3749, result)
}

func Test_Day7_Task1_InputData(t *testing.T) {
	result := day7.Task1("../day7/input.data")
	assert.Equal(t, 850435817339, result)
}

func Test_Day7_Task2_TestData(t *testing.T) {
	result := day7.Task2("../day7/test1.data")
	assert.Equal(t, 11387, result)
}

func Test_Day7_Task2_InputData(t *testing.T) {
	result := day7.Task2("../day7/input.data")
	assert.Equal(t, 31879108840545, result)
}
