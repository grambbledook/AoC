package util

import (
	"bufio"
	"os"
)

func ReadLinesFunc[T any](name string, f func(string) []T) [][]T {
	open, _ := os.Open(name)
	defer open.Close()

	in := bufio.NewReader(open)

	var lines [][]T
	for line, _, err := in.ReadLine(); err == nil; line, _, err = in.ReadLine() {
		lines = append(lines, f(string(line)))
	}

	return lines
}

func ReadLines(name string) []string {
	open, _ := os.Open(name)
	defer open.Close()

	in := bufio.NewReader(open)

	var lines []string
	for line, _, err := in.ReadLine(); err == nil; line, _, err = in.ReadLine() {
		lines = append(lines, string(line))
	}

	return lines
}

func ReadMatrix(name string) [][]rune {
	open, _ := os.Open(name)
	defer open.Close()

	in := bufio.NewReader(open)

	var lines [][]rune
	for line, _, err := in.ReadLine(); err == nil; line, _, err = in.ReadLine() {
		lines = append(lines, []rune(string(line)))
	}

	return lines
}
