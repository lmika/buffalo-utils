package slices

// ForEachWithError runs the passed in function for each element of T. If an error
// is encountered, the error is returned immediately and any subsequence elements
// will not be processed.
func ForEachWithError[T any](ts []T, fn func(t T) error) error {
	for _, t := range ts {
		if err := fn(t); err != nil {
			return err
		}
	}
	return nil
}
