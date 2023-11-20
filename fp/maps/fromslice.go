package maps

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
