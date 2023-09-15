package slices_test

import (
	"testing"

	"github.com/lmika/gopkgs/fp/slices"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("should return a mapped slice", func(t *testing.T) {
		ts := []int{1, 2, 3}
		us := slices.Map(ts, func(x int) int { return x + 2 })

		assert.Equal(t, []int{3, 4, 5}, us)
	})

	t.Run("should return nil if passed in nil", func(t *testing.T) {
		ts := []int(nil)
		us := slices.Map(ts, func(x int) int { return x + 2 })

		assert.Nil(t, us)
	})
}
