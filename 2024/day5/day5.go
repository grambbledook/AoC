package day5

import (
	"github.com/grambbledook/adventofcode2023/assert"
	"github.com/grambbledook/adventofcode2023/util"
	"iter"
	"strconv"
	"strings"
)

type Graph map[string][]string

func (g *Graph) iter() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for u, vs := range *g {
			for _, v := range vs {
				if !yield(u, v) {
					return
				}
			}
		}
	}
}

func Task1(file string) int {
	lines := util.ReadLines(file)
	graph, input := prepare(lines)

	var result int

	for _, line := range input {
		if !isValid(graph, line) {
			continue
		}

		mid := len(line) / 2
		val, err := strconv.Atoi(line[mid])

		assert.NoError(err)
		result += val
	}

	return result
}

func Task2(file string) int {
	lines := util.ReadLines(file)
	graph, input := prepare(lines)

	var result int

	for _, line := range input {
		if isValid(graph, line) {
			continue
		}

		fixed := topsort(graph, line)

		mid := len(fixed) / 2
		val, err := strconv.Atoi(fixed[mid])

		assert.NoError(err)
		result += val
	}

	return result
}

func prepare(lines []string) (Graph, [][]string) {
	var idx int

	graph := make(Graph)

	for idx = 0; idx < len(lines); idx++ {
		line := lines[idx]

		if line == "" {
			break
		}

		vals := strings.FieldsFunc(line, util.IsPipe)

		assert.Assert(len(vals) == 2)
		graph[vals[0]] = append(graph[vals[0]], vals[1])
	}

	input := make([][]string, 0)

	for idx = idx + 1; idx < len(lines); idx++ {
		vals := strings.FieldsFunc(lines[idx], util.IsComma)
		input = append(input, vals)
	}

	return graph, input
}

func isValid(graph Graph, vals []string) bool {
	positions := make(map[string]int)

	for idx, val := range vals {
		positions[val] = idx
	}

	for u, v := range graph.iter() {
		posU := util.GetOrDefault(positions, u, -1)
		posV := util.GetOrDefault(positions, v, -1)

		if posU == -1 || posV == -1 {
			continue
		}

		if posU > posV {
			return false
		}
	}
	return true
}

func topsort(graph Graph, vals []string) []string {
	indegree := make(map[string]int)

	for _, u := range vals {
		for _, v := range graph[u] {
			indegree[v] += 1
		}
	}

	result := make([]string, 0)

	queue := make([]string, 0)
	remainingNodes := make(map[string]bool)

	for _, val := range vals {
		if indegree[val] == 0 {
			queue = append(queue, val)
		}
		remainingNodes[val] = true
	}

	for len(queue) > 0 {
		lvl := queue
		queue = make([]string, 0)

		for _, u := range lvl {
			remainingNodes[u] = false
			result = append(result, u)

			for _, v := range graph[u] {
				if !remainingNodes[v] {
					continue
				}

				indegree[v] -= 1

				if indegree[v] == 0 {
					queue = append(queue, v)
				}
			}
		}
	}

	return result
}
