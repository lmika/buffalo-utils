package slices

func Contains[T comparable](ts []T, needle T) bool {
	for _, t := range ts {
		if t == needle {
			return true
		}
	}
	return false
}
