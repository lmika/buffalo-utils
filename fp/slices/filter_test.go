package slices_test

import (
	"strings"
	"testing"

	"github.com/lmika/gopkgs/fp/slices"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	var (
		ints = []int{1, 2, 3, 4, 5}
		strs = []string{"foo", "bar", "baz"}
	)

	t.Run("should filter items matching the predicate", func(t *testing.T) {
		assert.Equal(t, []int{2, 4}, slices.Filter(ints, func(x int) bool { return x%2 == 0 }))
		assert.Equal(t, []string{"bar", "baz"}, slices.Filter(strs, func(x string) bool { return strings.Contains(x, "b") }))
	})

	t.Run("should return nil if the passed in slice is nil", func(t *testing.T) {
		assert.Nil(t, slices.Filter(nil, func(x int) bool { return x%2 == 0 }))
	})

	t.Run("should return empty slice if the passed in slice is empty slice", func(t *testing.T) {
		assert.Equal(t, []int{}, slices.Filter([]int{}, func(x int) bool { return x%2 == 0 }))
	})
}
