package lang

import (
	"fmt"
)

type IdentityMap[T any] struct {
	Map map[string]T
}

func NewIdentityMap[T any]() IdentityMap[T] {
	return IdentityMap[T]{
		Map: make(map[string]T),
	}
}

func (m *IdentityMap[T]) Add(obj *T) {
	m.Map[fmt.Sprintf("%p", &obj)] = *obj
}

func (m *IdentityMap[T]) Values() []T {
	return Values(m.Map)
}
