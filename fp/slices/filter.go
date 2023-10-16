package slices

// Filter returns a slice containing all the elements of ts for which the passed in
// predicate returns true.  If no items match the predicate, the function will return
// an empty slice.  If ts is nil, the function will also return nil.
func Filter[T any](ts []T, predicate func(t T) bool) []T {
	if ts == nil {
		return nil
	}

	filteredTs := make([]T, 0)
	for _, t := range ts {
		if predicate(t) {
			filteredTs = append(filteredTs, t)
		}
	}
	return filteredTs
}
