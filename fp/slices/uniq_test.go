package slices_test

import (
	"fmt"
	"testing"

	"github.com/lmika/gopkgs/fp/slices"
	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	t.Run("should return a slice with unique elements", func(t *testing.T) {
		scenarios := []struct {
			in   []int
			want []int
		}{
			{in: nil, want: nil},
			{in: []int{}, want: []int{}},
			{in: []int{2}, want: []int{2}},
			{in: []int{1, 2}, want: []int{1, 2}},
			{in: []int{2, 2}, want: []int{2}},
			{in: []int{3, 1, 4, 2, 3, 5, 1, 4}, want: []int{3, 1, 4, 2, 5}},
		}

		for i, s := range scenarios {
			t.Run(fmt.Sprint(i), func(t *testing.T) {
				assert.Equal(t, s.want, slices.Uniq(s.in))
			})
		}
	})
}
