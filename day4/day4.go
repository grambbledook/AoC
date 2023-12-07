package day4

import (
	"github.com/grambbledook/adventofcode2023/lang"
	"math"
	"strings"
)

func Task1(str []string) int {
	return lang.Fold(str, 0, func(a int, b string) int {
		_, total := processCard(b)

		if total == 0 {
			return a
		}

		return int(math.Pow(2, float64(total-1)))
	})
}

func Task2(data []string) int {
	cards := make(map[int][]int, len(data))

	for _, c := range data {
		card, total := processCard(c)
		cards[card] = lang.IntRange(card+1, card+total+1)
	}

	total := 0
	for card := range cards {
		total += dfs(cards, card)
	}

	return total
}

func dfs(graph map[int][]int, u int) int {
	total := 1
	for _, v := range graph[u] {
		total += dfs(graph, v)
	}
	return total
}

func processCard(str string) (int, int) {
	parsed := strings.Fields(str)

	i := 2
	winningNumbers := make(map[int]bool)
	for ; i < len(parsed); i++ {
		if parsed[i] == "|" {
			break
		}
		winningNumbers[lang.Int(parsed[i])] = true
	}

	yourNumbers := make(map[int]bool)
	for ; i < len(parsed); i++ {
		yourNumbers[lang.Int(parsed[i])] = true
	}

	total := 0
	for number := range yourNumbers {
		if winningNumbers[number] {
			total += 1
		}
	}

	return lang.Int(strings.TrimRight(parsed[1], ":")), total
}
