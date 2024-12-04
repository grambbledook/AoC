package day4

import (
	"github.com/grambbledook/adventofcode2023/util"
)

const SearchWord = "XMAS"

func Task1(file string) int {
	lines := util.ReadLines(file)

	n, m := len(lines), len(lines[0])

	type Key struct {
		x int
		y int
	}

	var result int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if lines[i][j] != 'X' {
				continue
			}

			idx := 0
			for k := 0; k < min(n, i+4); k++ {
				if lines[i+k][j] != SearchWord[idx] {
					break
				}
				idx++
			}
			result += util.FromBool(idx == len(SearchWord))

		}
	}

	return result
}
