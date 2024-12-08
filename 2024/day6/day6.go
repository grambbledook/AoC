package day6

import (
	"github.com/grambbledook/adventofcode2023/util"
)

type Grid []string

const (
	Up    int = iota
	Right int = iota
	Down  int = iota
	Left  int = iota
)

func (g *Grid) oob(x, y int) bool {
	grid := *g

	if x < 0 || y < 0 {
		return true
	}

	if x >= len(grid) || y >= len(grid[0]) {
		return true
	}

	return false
}

func Task1(file string) int {
	lines := util.ReadLines(file)
	var result int

	for i, line := range lines {
		for j, char := range line {
			if char != '^' {
				continue
			}
			result += dfs(Up, i, j, lines, make(map[util.Point]int))
		}
	}

	return result
}

func Task2(file string) int {
	return Task1(file)
}

func dfs(dir int, i, j int, grid Grid, visited map[util.Point]int) int {
	if grid.oob(i, j) {
		return 0
	}

	if grid[i][j] == '#' {
		return -1
	}

	result := 0

	mask := 1 << dir
	cell := util.Point{X: i, Y: j}

	if visited[cell] == 0 {
		result += 1
	}

	if visited[cell]&mask != 0 {
		return 0
	}

	visited[cell] |= mask

	for k := 0; k < 4; k++ {
		nextDir := (dir + k) % 4
		di, dj := step(i, j, nextDir)
		res := dfs(nextDir, di, dj, grid, visited)

		if res != -1 {
			return res + result
		}
	}

	return result
}

func step(x, y, dir int) (int, int) {
	switch dir {
	case Up:
		return x + -1, y + 0
	case Down:
		return x + 1, y + 0
	case Left:
		return x + 0, y + -1
	default:
		return x + 0, y + 1
	}
}
