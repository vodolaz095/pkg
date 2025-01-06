package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequence(t *testing.T) {
	t.Run("int", func(tt *testing.T) {
		seed := func(i int) int {
			return i + 1
		}
		val := Sequence[int](0, seed)
		assert.Equal(tt, 0, len(val))

		val = Sequence[int](2, seed)
		assert.Equal(tt, 2, len(val))
		assert.Equal(tt, 1, val[0])
		assert.Equal(tt, 2, val[1])

		val = Sequence[int](3, seed)
		assert.Equal(tt, 3, len(val))
		assert.Equal(tt, 1, val[0])
		assert.Equal(tt, 2, val[1])
		assert.Equal(tt, 3, val[2])
	})
}
