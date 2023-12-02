package util

import (
	"bufio"
	"os"
)

func ReadFile(name string) []string {
	open, _ := os.Open(name)
	defer open.Close()

	in := bufio.NewReader(open)

	var lines []string
	for line, _, err := in.ReadLine(); err == nil; line, _, err = in.ReadLine() {
		lines = append(lines, string(line))
	}

	return lines
}
