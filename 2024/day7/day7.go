package day7

import (
	"github.com/grambbledook/adventofcode2023/util"
	"strings"
)

type Task struct {
	vals     []int
	expected int
	concat   bool
}

func Task1(file string) int {
	lines := util.ReadLines(file)

	var result int

	for _, line := range lines {
		spla := strings.Split(line, ":")

		task := Task{
			vals:     util.Ints(spla[1]),
			expected: util.Int(spla[0]),
			concat:   false,
		}

		if task.isValid(1, 0) {
			result += task.expected
		}
	}

	return result
}

func Task2(file string) int {
	lines := util.ReadLines(file)

	var result int

	for _, line := range lines {
		spla := strings.Split(line, ":")

		task := Task{
			vals:     util.Ints(spla[1]),
			expected: util.Int(spla[0]),
			concat:   true,
		}

		if task.isValid(1, 0) {
			result += task.expected
		}

		println(result)
	}

	return result
}

func (t *Task) isValid(idx int, total int) bool {
	if total > t.expected {
		return false
	}

	if idx == len(t.vals) {
		return t.expected == total
	}

	var mul, add, con bool

	if idx == 1 {
		mul = t.isValid(idx+1, t.vals[0]*t.vals[1])
		add = t.isValid(idx+1, t.vals[0]+t.vals[1])
		con = t.isValid(idx+1, concat(t.vals[0], t.vals[1]))
	} else {
		mul = t.isValid(idx+1, total*t.vals[idx])
		add = t.isValid(idx+1, total+t.vals[idx])
		con = t.isValid(idx+1, concat(total, t.vals[idx]))
	}

	return add || mul || (con && t.concat)
}

func concat(a, b int) int {
	for b != 0 {
		rem, mod := b/10, b%10
		a, b = a*10+mod, rem
	}
	return a
}
