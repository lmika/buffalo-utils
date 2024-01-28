package maps_test

import (
	"testing"

	"github.com/lmika/gopkgs/fp/maps"
	"github.com/stretchr/testify/assert"
)

func TestToSlice(t *testing.T) {
	type pair struct {
		left  int
		right string
	}

	ms := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	pairs := maps.ToSlice(ms, func(k int, v string) pair { return pair{k, v} })

	assert.Len(t, pairs, 3)
	assert.Contains(t, pairs, pair{1, "one"})
	assert.Contains(t, pairs, pair{2, "two"})
	assert.Contains(t, pairs, pair{3, "three"})
}
