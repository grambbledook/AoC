package lang

import (
	"regexp"
	"strconv"
	"strings"
)

func Int(s string) int {
	if num, err := strconv.Atoi(s); err == nil {
		return num
	} else {
		return 0
	}
}

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func Fields(s string) func(int, bool) string {
	fields := strings.Fields(s)

	return func(idx int, strip bool) string {
		if strip == false {
			return fields[idx]
		}

		return clearString(fields[idx])
	}
}
