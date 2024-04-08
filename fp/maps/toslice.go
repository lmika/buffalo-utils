package maps

func ToSlice[K comparable, V, T any](m map[K]V, fn func(k K, v V) T) []T {
	if m == nil {
		return nil
	}

	ts := make([]T, 0, len(m))
	for k, v := range m {
		ts = append(ts, fn(k, v))
	}
	return ts
}

func ToSliceWithError[K comparable, V, T any](m map[K]V, fn func(k K, v V) (T, error)) ([]T, error) {
	if m == nil {
		return nil, nil
	}

	ts := make([]T, 0, len(m))
	for k, v := range m {
		w, err := fn(k, v)
		if err != nil {
			return nil, err
		}
		ts = append(ts, w)
	}
	return ts, nil
}
