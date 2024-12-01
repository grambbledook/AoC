package test

import (
	"github.com/grambbledook/adventofcode2023/2023/day5"
	"github.com/grambbledook/adventofcode2023/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Day5_1_TestData(t *testing.T) {
	result := day5.Task1(util.ReadLines("../day5/test"))
	assert.Equal(t, 35, result)
}

func Test_Day5_1_TaskData(t *testing.T) {
	result := day5.Task1(util.ReadLines("../day5/input"))
	assert.Equal(t, 165788812, result)
}

func Test_Day5_2_TestData(t *testing.T) {
	result := day5.Task2(util.ReadLines("../day5/test"))
	assert.Equal(t, 46, result)
}

func Test_Day5_2_TaskData(t *testing.T) {
	result := day5.Task2(util.ReadLines("../day5/input"))
	assert.Equal(t, 1928058, result)
}
