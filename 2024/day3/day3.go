package day3

import (
	"github.com/grambbledook/adventofcode2023/util"
	"strings"
	"unicode"
)

type Instruction interface {
	write(rune) bool
	ignore() bool
}

type Virgin struct {
}

func (s *Virgin) write(_ rune) bool {
	return false
}

func (s *Virgin) ignore() bool {
	return false
}

type Chad struct {
	last    rune
	done    bool
	valid   bool
	ignores bool
}

func NewInstruction(chad bool) Instruction {
	if chad {
		return &Chad{done: false, valid: true, ignores: false}
	}

	return &Virgin{}
}

func (s *Chad) write(c rune) bool {
	if s.done {
		return false
	}

	if !s.valid {
		return false
	}

	switch {
	case c == 'd':
		s.valid = s.last == 0
	case c == 'o':
		s.valid = s.last == 'd'
	case c == 'n':
		s.valid = s.last == 'o'
	case c == '\'':
		s.valid = s.last == 'n'
	case c == 't':
		s.valid = s.last == '\''
	case c == '(':
		s.valid = s.last == 't' || s.last == 'o'
		s.ignores = s.last == 't'
	case c == ')':
		s.valid = s.last == '('
		s.done = true
	default:
		s.valid = false
	}

	s.last = c
	return s.done
}

func (s *Chad) ignore() bool {
	if !s.valid {
		return false
	}

	return s.done && s.ignores
}

type Multiplier struct {
	a     int
	b     int
	last  rune
	valid bool
	done  bool
}

func NewMultiplier() Multiplier {
	return Multiplier{a: 0, b: 0, last: 0, valid: true}
}

func (s *Multiplier) write(c rune) int {
	if !s.valid {
		return 0
	}

	if s.done {
		return 0
	}

	result := 0
	switch {
	case c == 'm':
		s.valid = s.last == 0
	case c == 'u':
		s.valid = s.last == 'm'
	case c == 'l':
		s.valid = s.last == 'u'
	case c == '(':
		s.valid = s.last == 'l'
	case c == ',':
		s.a, s.b = s.b, 0
	case c == ')' && s.valid:
		result += s.b * s.a
		s.done = true
	case unicode.IsDigit(c):
		s.b = s.b*10 + int(c-'0')
	default:
		s.valid = false
	}

	s.last = c
	return result
}

func Task1(file string) int {
	lines := util.ReadLines(file)

	var result int

	for _, line := range lines {
		val := Parse(line, false)
		result += val
	}

	return result
}

func Task2(file string) int {
	lines := util.ReadLines(file)

	result := Parse(strings.Join(lines, ""), true)

	return result
}

func Parse(line string, chad bool) int {
	result := 0
	state := NewMultiplier()

	next := NewInstruction(chad)
	instruction := NewInstruction(false)

	for _, c := range line {
		if c == 'd' {
			next = NewInstruction(chad)
		}

		if done := next.write(c); done {
			instruction = next
		}

		if instruction.ignore() {
			continue
		}

		if c == 'm' {
			state = NewMultiplier()
		}

		result += state.write(c)
	}

	return result
}
