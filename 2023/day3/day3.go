package day3

import (
	lang2 "github.com/grambbledook/adventofcode2023/2023/lang"
	"github.com/grambbledook/adventofcode2023/2023/task"
	"github.com/grambbledook/adventofcode2023/util"
	"unicode"
)

var directions = [][]int{
	[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}, []int{1, 1}, []int{-1, -1}, []int{1, -1}, []int{-1, 1},
}

func Compute(f task.MatrixTask, file string) int {
	grid := util.ReadLines(file)
	return f(grid)
}

func Task1(grid []string) int {
	s := parse(grid)
	print(s)
	p := parts(s)
	print(p)
	maped := lang2.Map(p, func(s Serial) int { return s.Num })
	val := util.Sum(maped)
	return val
}

func Task2(grid []string) int {
	return 0
}

func parse(grid []string) [][]any {
	parsed := make([][]any, len(grid))
	for i := range parsed {
		parsed[i] = make([]any, len(grid[i]))
	}

	for i, line := range grid {
		num := ""

		for j, c := range line + "." {
			if unicode.IsDigit(c) {
				num += string(c)
			} else if num != "" {
				serial := Serial{lang2.Int(num)}
				for k := 0; k < len(num); k++ {
					parsed[i][j-k] = serial
				}
				num = ""
			}

			if c != '.' && !unicode.IsDigit(c) {
				parsed[i][j] = Symbol{c}
			}
		}
	}

	return parsed
}

func parts(grid [][]any) []Serial {
	parts := lang2.NewIdentityMap[Serial]()

	for i, line := range grid {
		for j, c := range line {
			switch c.(type) {
			case Symbol:
				serials := scan(grid, i, j)
				for i := range serials {
					serial := serials[i]
					parts.Add(&serial)
				}
			}
		}
	}

	return parts.Values()
}

func scan(grid [][]any, i int, j int) []Serial {
	parts := make([]Serial, 0)
	for _, dir := range directions {
		x, y := i+dir[0], j+dir[1]
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
			continue
		}

		switch grid[x][y].(type) {
		case Serial:
			parts = append(parts, grid[x][y].(Serial))
		}
	}
	return parts
}

type Serial struct {
	Num int
}

type Symbol struct {
	Char rune
}
