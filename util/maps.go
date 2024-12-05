package util

func GetOrDefault[K comparable, V any](dict map[K]V, key K, def V) V {
	val, ok := dict[key]
	if !ok {
		return def
	}
	return val
}
