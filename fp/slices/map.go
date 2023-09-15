package slices

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
