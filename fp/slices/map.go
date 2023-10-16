package slices

// Map returns a new slice containing the elements of ts transformed by the passed in function.
func Map[T, U any](ts []T, fn func(t T) U) (us []U) {
	if ts == nil {
		return nil
	}

	us = make([]U, len(ts))
	for i, t := range ts {
		us[i] = fn(t)
	}
	return us
}

// Map returns a new slice containing the elements of ts transformed by the passed in function, which
// can either either a mapped value of U, or an error. If the mapping function returns an error, MapWithError
// will return nil and the returned error.
func MapWithError[T, U any](ts []T, fn func(t T) (U, error)) (us []U, err error) {
	if ts == nil {
		return nil, nil
	}

	us = make([]U, len(ts))
	for i, t := range ts {
		var e error
		us[i], e = fn(t)
		if e != nil {
			return nil, e
		}
	}
	return us, nil
}
