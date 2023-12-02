package lang

func Values[K string | int, V interface{}](dict map[K]V) []V {
	var values []V
	for _, value := range dict {
		values = append(values, value)
	}
	return values
}

func Map[A, B interface{}](arr []A, f func(A) B) []B {
	var mapped []B
	for _, a := range arr {
		mapped = append(mapped, f(a))
	}
	return mapped
}

func Fold[A, B interface{}](arr []B, identity A, f func(A, B) A) A {
	acc := identity
	for _, a := range arr {
		acc = f(acc, a)
	}
	return acc
}
