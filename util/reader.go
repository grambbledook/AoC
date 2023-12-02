package util

import (
	"bufio"
	"os"
)

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
