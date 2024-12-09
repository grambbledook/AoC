package test

import (
	"github.com/grambbledook/adventofcode2023/2024/day3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day3_Task1_TestData(t *testing.T) {
	result := day3.Task1("../day3/test1.data")
	assert.Equal(t, 161, result)
}

func Test_Day3_Task1_InputData(t *testing.T) {
	result := day3.Task1("../day3/input.data")
	assert.Equal(t, 173785482, result)
}

func Test_Day3_Task2_TestData(t *testing.T) {
	result := day3.Task2("../day3/test2.data")
	assert.Equal(t, 48, result)
}

func Test_Day3_Task2_InputData(t *testing.T) {
	result := day3.Task2("../day3/input.data")
	assert.Equal(t, 83158140, result)
}

func Test_Day3_Util_Func_Virgin(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"mul(44,46)", 2024},
		{"mul(123,4)", 492},
		{"mul(123,4))", 492},
		{"do()mul(123,4)", 492},
		{"don't()mul(123,4)", 492},
		{"mul(4*", 0},
		{"mul(6,9!", 0},
		{"?(12,34)", 0},
		{"mul ( 2 , 4 )", 0},
	}
	for _, test := range tests {
		result := day3.Parse(test.input, false)
		if result != test.expected {
			t.Errorf("Test failed. Input: %v, Expected: %v, Result: %v", test.input, test.expected, result)
		}
	}
}

func Test_Day3_Util_Func_Chad(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"mul(44,46)", 2024},
		{"mul(123,4)", 492},
		{"do()mul(123,4)", 492},
		{"don't()mul(123,4)", 0},
		{"don't()mul(123,4)mul(1,5)", 0},
		{"don't()_do()_mul(123,4)", 492},
		{"don't()\nmul(123,4)mul(1,5)", 0},
		{"mul(4*", 0},
		{"mul(6,9!", 0},
		{"?(12,34)", 0},
		{"mul ( 2 , 4 )", 0},
	}
	for _, test := range tests {
		result := day3.Parse(test.input, true)
		if result != test.expected {
			t.Errorf("Test failed. Input: %v, Expected: %v, Result: %v", test.input, test.expected, result)
		}
	}
}
