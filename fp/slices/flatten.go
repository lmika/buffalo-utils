package slices

func Flattern[T any](tss [][]T) []T {
	if len(tss) == 0 {
		return nil
	}

	entireLen := 0
	for _, ts := range tss {
		entireLen += len(ts)
	}

	newTs := make([]T, 0, entireLen)
	for _, ts := range tss {
		newTs = append(newTs, ts...)
	}

	return newTs
}
