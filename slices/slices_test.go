package slices_test

import (
	"testing"

	"github.com/kristofaranyos/go-utils/slices"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Run("filter odds", func(t *testing.T) {
		arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		out := slices.Filter(arr, func(t int) bool {
			return t%2 == 0
		})

		assert.Equal(t, []int{0, 2, 4, 6, 8, 10}, out)
	})

	t.Run("filter for one element outside closure", func(t *testing.T) {
		arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		localTwo := 2

		out := slices.Filter(arr, func(t int) bool {
			return t == localTwo
		})

		assert.Equal(t, []int{2}, out)
	})
}
