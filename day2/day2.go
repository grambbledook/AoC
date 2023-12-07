package day2

import (
	"github.com/grambbledook/adventofcode2023/lang"
	"github.com/grambbledook/adventofcode2023/task"
	"github.com/grambbledook/adventofcode2023/util"
	"strings"
)

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Compute(f task.LineTask, file string) int {
	lines := util.ReadLines(file)

	return lang.Fold(lines, 0, func(a int, b string) int { return a + f(b) })
}

func Task1(str string) int {
	parsed := strings.Split(str, ":")

	for _, seed := range strings.Split(parsed[1], ";") {
		for _, dice := range strings.Split(seed, ",") {
			fields := lang.Fields(dice)

			count := lang.Int(fields(0, true))
			color := fields(1, true)

			if limits[color] < count {
				return 0
			}
		}
	}

	return lang.Int(strings.Fields(parsed[0])[1])
}

func Task2(str string) int {
	parsed := strings.Split(str, ":")

	counts := make(map[string]int)

	for _, seed := range strings.Split(parsed[1], ";") {
		for _, dice := range strings.Split(seed, ",") {
			fields := lang.Fields(dice)

			color := fields(1, true)
			count := lang.Int(fields(0, true))

			counts[color] = util.Max(count, counts[color])
		}
	}

	return util.Product(lang.Values(counts))
}
