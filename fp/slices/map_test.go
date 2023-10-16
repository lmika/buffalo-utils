package slices_test

import (
	"errors"
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

func TestMapWithError(t *testing.T) {
	t.Run("should return a mapped slice with no error", func(t *testing.T) {
		ts := []int{1, 2, 3}

		us, err := slices.MapWithError(ts, func(x int) (int, error) { return x + 2, nil })

		assert.Equal(t, []int{3, 4, 5}, us)
		assert.NoError(t, err)
	})

	t.Run("should return nil with an error when mapping function returns an error", func(t *testing.T) {
		ts := []int{1, 2, 3}

		us, err := slices.MapWithError(ts, func(x int) (int, error) {
			if x == 2 {
				return 0, errors.New("bang")
			}
			return x + 2, nil
		})

		assert.Nil(t, us)
		assert.Error(t, err)
	})

	t.Run("should return nil if passed in nil", func(t *testing.T) {
		ts := []int(nil)
		us, err := slices.MapWithError(ts, func(x int) (int, error) { return x + 2, nil })

		assert.Nil(t, us)
		assert.NoError(t, err)
	})
}
