package day4

import (
	"github.com/grambbledook/adventofcode2023/util"
	"iter"
)

const SearchWord = "XMAS"

type Grid struct {
	n int
	m int
}

func (t *Grid) oob(i, j int) bool {
	return !t.noob(i, j)
}

func (t *Grid) noob(i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}

	if i >= t.n || j >= t.m {
		return false
	}

	return true
}

func (t *Grid) iter() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i := 0; i < t.n; i++ {
			for j := 0; j < t.m; j++ {
				if !yield(i, j) {
					return
				}
			}
		}
	}
}

func Task1(file string) int {
	lines := util.ReadLines(file)

	var result int
	grid := Grid{len(lines), len(lines[0])}

	for i, j := range grid.iter() {
		if lines[i][j] != 'X' {
			continue
		}

		counts := [8]int{}
		for k := 0; k < len(SearchWord); k++ {
			next := SearchWord[k]

			if grid.noob(i+k, j) && lines[i+k][j] == next {
				counts[0]++
			}

			if grid.noob(i, j+k) && lines[i][j+k] == next {
				counts[1]++
			}

			if grid.noob(i+k, j+k) && lines[i+k][j+k] == next {
				counts[2]++
			}

			if grid.noob(i-k, j+k) && lines[i-k][j+k] == next {
				counts[3]++
			}

			if grid.noob(i+k, j-k) && lines[i+k][j-k] == next {
				counts[4]++
			}

			if grid.noob(i-k, j-k) && lines[i-k][j-k] == next {
				counts[5]++
			}

			if grid.noob(i-k, j) && lines[i-k][j] == next {
				counts[6]++
			}

			if grid.noob(i, j-k) && lines[i][j-k] == next {
				counts[7]++
			}
		}

		for _, count := range counts {
			if count != len(SearchWord) {
				continue
			}
			result += 1
		}
	}

	return result
}

func Task2(file string) int {
	lines := util.ReadLines(file)

	var result int
	grid := Grid{len(lines), len(lines[0])}

	for i, j := range grid.iter() {
		if lines[i][j] != 'A' {
			continue
		}

		if grid.oob(i-1, j-1) || grid.oob(i+1, j+1) {
			continue
		}

		if grid.oob(i+1, j-1) || grid.oob(i-1, j+1) {
			continue
		}

		if lines[i-1][j-1] == lines[i+1][j+1] {
			continue
		}

		if lines[i+1][j-1] == lines[i-1][j+1] {
			continue
		}

		counts := [26]int{}

		counts[lines[i-1][j-1]-'A'] += 1
		counts[lines[i+1][j+1]-'A'] += 1
		counts[lines[i-1][j+1]-'A'] += 1
		counts[lines[i+1][j-1]-'A'] += 1

		if counts['M'-'A'] == 2 && counts['S'-'A'] == 2 {
			result += 1
		}
	}

	return result
}
