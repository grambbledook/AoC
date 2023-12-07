package lang

func Values[K comparable, V any](dict map[K]V) []V {
	var values []V
	for _, value := range dict {
		values = append(values, value)
	}
	return values
}

func FilterKeys[M ~map[K]V, K comparable, V any](dict M, f func(K) bool) M {
	var filtered = make(M)
	for key, value := range dict {
		if f(key) {
			filtered[key] = value
		}
	}
	return filtered
}

func FilterValues[M ~map[K]V, K comparable, V any](dict M, f func(V) bool) M {
	var filtered = make(M)
	for key, value := range dict {
		if f(value) {
			filtered[key] = value
		}
	}
	return filtered
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

func IntRange(start int, end int) []int {
	var rng []int
	for i := start; i < end; i++ {
		rng = append(rng, i)
	}
	return rng
}
