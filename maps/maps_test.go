package maps_test

import (
	"testing"

	"go-utils/maps"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	m := map[string]string{
		"England": "London",
		"Germany": "Berlin",
		"France":  "Paris",
	}

	t.Run("contains Berlin", func(t *testing.T) {
		assert.Equal(t, true, maps.Contains(m, "Berlin"))
	})

	t.Run("doesn't contain Brussels", func(t *testing.T) {
		assert.Equal(t, false, maps.Contains(m, "Brussels"))
	})
}

func TestContainsKey(t *testing.T) {
	m := map[string]string{
		"England": "London",
		"Germany": "Berlin",
		"France":  "Paris",
	}

	t.Run("contains Germany", func(t *testing.T) {
		assert.Equal(t, true, maps.ContainsKey(m, "Germany"))
	})

	t.Run("doesn't contain Belgium", func(t *testing.T) {
		assert.Equal(t, false, maps.ContainsKey(m, "Belgium"))
	})
}

func TestClear(t *testing.T) {
	m := map[string]string{
		"England": "London",
		"Germany": "Berlin",
		"France":  "Paris",
	}

	maps.Clear(m)

	assert.Len(t, m, 0)
}
