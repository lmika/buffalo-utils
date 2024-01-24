package slices

func Uniq[T comparable](ts []T) []T {
	if len(ts) < 2 {
		return ts
	}

	outT := make([]T, 0)
	seenT := make(map[T]struct{})

	for _, t := range ts {
		if _, ok := seenT[t]; !ok {
			outT = append(outT, t)
			seenT[t] = struct{}{}
		}
	}

	return outT
}
