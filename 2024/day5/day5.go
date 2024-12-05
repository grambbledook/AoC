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

	return process(graph, input, false)
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

func process(graph Graph, input [][]string, fix bool) int {
	var result int

	for _, line := range input {
		valid := isValid(graph, line)

		if !valid && !fix {
			continue
		}

		if !valid {
			line = topsort(graph, line)
		}

		mid := len(line) / 2
		val, err := strconv.Atoi(line[mid])
		assert.NoError(err)

		result += val
	}

	return result
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
	return vals
}
