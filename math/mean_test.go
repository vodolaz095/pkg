package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMean(t *testing.T) {
	assert.Equal(t, 0, Mean[int]())
	assert.Equal(t, 5, Mean[int](2, 8))
	assert.Equal(t, .33, Mean[float64](0.34, .32, .33))
}
