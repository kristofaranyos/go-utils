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

func TestCast(t *testing.T) {
	t.Run("cast to float", func(t *testing.T) {
		arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		out := slices.Cast(arr, func(t int) float32 {
			return float32(t)
		})

		assert.Equal(t, []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, out)
	})

	t.Run("cast to custom type", func(t *testing.T) {
		type customType string

		arr := []string{"value1", "value2"}

		out := slices.Cast(arr, func(t string) customType {
			return customType(t)
		})

		assert.Equal(t, []customType{"value1", "value2"}, out)
	})
}
