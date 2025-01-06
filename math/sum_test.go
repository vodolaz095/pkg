package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 0, Sum[int]())
	assert.Equal(t, 1, Sum[int](1))
	assert.Equal(t, 3, Sum[int](1, 2))
	assert.Equal(t, -1, Sum[int](1, -2))
	assert.Equal(t, float64(1), Sum[float64](0.34, .32, .33, .01))
}
