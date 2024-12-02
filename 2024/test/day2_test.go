package test

import (
	"github.com/grambbledook/adventofcode2023/2024/day2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day2_Task1_TestData(t *testing.T) {
	result := day2.Task1("../day2/test1")
	assert.Equal(t, 2, result)
}

func Test_Day2_Task1_InputData(t *testing.T) {
	result := day2.Task1("../day2/input")
	assert.Equal(t, 631, result)
}

func Test_Day2_Task2_TestData(t *testing.T) {
	result := day2.Task2("../day2/test1")
	assert.Equal(t, 4, result)
}

func Test_Day2_Task2_InputData(t *testing.T) {
	result := day2.Task2("../day2/input")
	assert.Equal(t, 665, result)
}

func Test_Util_Func(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, true},
		{[]int{10, 9, 8, 7}, true},
		{[]int{1, 3, 6, 9}, true},
		{[]int{6, 2, 1}, false},
		{[]int{6, 6, 5}, false},
		{[]int{3, 4, 6}, true},
		{[]int{10, 7, 5, 4}, true},
	}
	for _, test := range tests {
		result := day2.Test(test.input)
		if result != test.expected {
			t.Errorf("Test failed. Input: %v, Expected: %v, Result: %v", test.input, test.expected, result)
		}
	}
}
