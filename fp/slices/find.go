package slices

func Contains[T comparable](ts []T, needle T) bool {
	for _, t := range ts {
		if t == needle {
			return true
		}
	}
	return false
}

func FindWhere[T comparable](ts []T, predicate func(t T) bool) (T, bool) {
	var zeroT T

	for _, t := range ts {
		if predicate(t) {
			return t, true
		}
	}
	return zeroT, false
}
