package slices_test

import (
	"testing"

	"github.com/lmika/gopkgs/fp/slices"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	var (
		ints = []int{1, 2, 3}
		strs = []string{"a", "b", "c"}
	)

	t.Run("should find items in the slice", func(t *testing.T) {
		assert.True(t, slices.Contains(ints, 2))
		assert.True(t, slices.Contains(strs, "c"))
	})

	t.Run("should return false if items not in slice", func(t *testing.T) {
		assert.False(t, slices.Contains(ints, 131))
		assert.False(t, slices.Contains(strs, "bla"))
	})
}
