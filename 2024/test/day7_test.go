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
	assert.Equal(t, 104824810233437, result)
}

func Test_Day7_Concat(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{12, 34, 1234},
		{56, 78, 5678},
		{9, 100, 9100},
		{45, 6, 456},
		{7, 89, 789},
	}

	for _, test := range testCases {
		result := day7.Concat(test.a, test.b)
		if result != test.expected {
			t.Errorf("Concat(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}
