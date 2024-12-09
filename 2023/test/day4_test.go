package test

import (
	"github.com/grambbledook/adventofcode2023/2023/day4"
	"github.com/grambbledook/adventofcode2023/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day4_1_TestData(t *testing.T) {
	result := day4.Task1(util.ReadLines("../day4/test1.data"))
	assert.Equal(t, 13, result)
}

func Test_Day4_1_TaskData(t *testing.T) {
	result := day4.Task1(util.ReadLines("../day4/input.data"))
	assert.Equal(t, 21568, result)
}

func Test_Day4_2_TestData(t *testing.T) {
	result := day4.Task2(util.ReadLines("../day4/test1.data"))
	assert.Equal(t, 30, result)
}

func Test_Day4_2_TaskData(t *testing.T) {
	result := day4.Task2(util.ReadLines("../day4/input.data"))
	assert.Equal(t, 11827296, result)
}
