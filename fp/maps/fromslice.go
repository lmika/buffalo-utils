package maps

func FromSliceGroups[T any, K comparable](ts []T, fn func(t T) K) map[K][]T {
	kvs := make(map[K][]T)
	for _, t := range ts {
		k := fn(t)
		kvs[k] = append(kvs[k], t)
	}
	return kvs
}

func FromSlice[T any, K comparable, V any](ts []T, fn func(t T) (K, V)) map[K]V {
	m, _ := FromSliceWithError(ts, func(t T) (k K, v V, _ error) {
		k, v = fn(t)
		return k, v, nil
	})
	return m
}

func FromSliceWithError[T any, K comparable, V any](ts []T, fn func(t T) (K, V, error)) (map[K]V, error) {
	kvs := make(map[K]V)
	for _, t := range ts {
		k, v, err := fn(t)
		if err != nil {
			return nil, err
		}
		kvs[k] = v
	}
	return kvs, nil
}
