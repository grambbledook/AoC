package day1

import (
	"github.com/grambbledook/adventofcode2023/lang"
	"github.com/grambbledook/adventofcode2023/task"
	"github.com/grambbledook/adventofcode2023/util"
	"maps"
	"math"
	"strings"
)

var wordsToDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var digitsToDigits = map[string]int{
	"1":    1,
	"2":    2,
	"3":    3,
	"4":    4,
	"5":    5,
	"6":    6,
	"7":    7,
	"8":    8,
	"9":    9,
	"zero": 0,
	"0":    0,
}

var taskOneDict map[string]int

var taskTwoDict map[string]int

func init() {
	taskOneDict = make(map[string]int, len(digitsToDigits))
	maps.Copy(taskOneDict, digitsToDigits)

	taskTwoDict = make(map[string]int, len(digitsToDigits)+len(wordsToDigits))
	maps.Copy(taskTwoDict, digitsToDigits)
	maps.Copy(taskTwoDict, wordsToDigits)
}

func Compute(f task.LineTask, file string) int {
	lines := util.ReadLines(file)

	return lang.Fold(lines, 0, func(a int, b string) int { return a + f(b) })
}

func Task1(line string) int {
	return process(taskOneDict, line)
}

func Task2(line string) int {
	return process(taskTwoDict, line)
}

func process(codes map[string]int, line string) int {
	a, apos := 0, math.MaxInt
	b, bpos := 0, math.MinInt

	for k, v := range codes {
		l := strings.Index(line, k)
		if l != -1 && l < apos {
			a, apos = v, l
		}

		r := strings.LastIndex(line, k)
		if r != -1 && r > bpos {
			b, bpos = v, r
		}
	}
	return a*10 + b
}
