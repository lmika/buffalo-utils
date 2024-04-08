package maps

func Values[K comparable, V any](m map[K]V) []V {
	vs := make([]V, 0, len(m))
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

func MapValues[K comparable, V any, W any](m map[K]V, fn func(v V) W) map[K]W {
	ws := make(map[K]W)
	for k, v := range m {
		ws[k] = fn(v)
	}
	return ws
}
