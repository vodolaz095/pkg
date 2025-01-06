package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseMinMax[T Number] struct {
	input []T
	ret   T
	index int
}

func validateMaxTestCase[T Number](t *testing.T, testCases []testCaseMinMax[T]) {
	var (
		returned      T
		indexReturned int
	)
	for i := range testCases {
		returned, indexReturned = Max[T](testCases[i].input...)
		assert.Equal(t, indexReturned, testCases[i].index,
			"test case %v (%v) - wrong index", i, testCases[i].input)
		assert.Equal(t, returned, testCases[i].ret,
			"test case %v (%v) - wrong return", i, testCases[i].input)
	}
}

func validateMinTestCase[T Number](t *testing.T, testCases []testCaseMinMax[T]) {
	var (
		returned      T
		indexReturned int
	)
	for i := range testCases {
		returned, indexReturned = Min[T](testCases[i].input...)
		assert.Equal(t, indexReturned, testCases[i].index,
			"test case %v (%v) - wrong index", i, testCases[i].input)
		assert.Equal(t, returned, testCases[i].ret,
			"test case %v (%v) - wrong return", i, testCases[i].input)
	}
}

func TestMax(t *testing.T) {
	t.Parallel()
	t.Run("uint", func(tt *testing.T) {
		testCases := []testCaseMinMax[uint]{
			{make([]uint, 0), 0, -1},
			{[]uint{1}, 1, 0},
			{[]uint{0}, 0, 0},
			{[]uint{1, 0}, 1, 0},
			{[]uint{0, 1}, 1, 1},
			{[]uint{0, 1, 2, 3}, 3, 3},
		}
		validateMaxTestCase[uint](tt, testCases)
	})
	t.Run("int", func(tt *testing.T) {
		testCases := []testCaseMinMax[int]{
			{make([]int, 0), 0, -1},
			{[]int{1}, 1, 0},
			{[]int{0}, 0, 0},
			{[]int{1, 0}, 1, 0},
			{[]int{0, 1}, 1, 1},
			{[]int{0, 1, 2, 3}, 3, 3},
			{[]int{0, 1, 2, -3}, 2, 2},
		}
		validateMaxTestCase[int](tt, testCases)
	})
	t.Run("float32", func(tt *testing.T) {
		testCases := []testCaseMinMax[float32]{
			{make([]float32, 0), 0, -1},
			{[]float32{1}, 1, 0},
			{[]float32{0}, 0, 0},
			{[]float32{1, 0}, 1, 0},
			{[]float32{0, 1}, 1, 1},
			{[]float32{0, 1, 2, 3}, 3, 3},
			{[]float32{0, 1, 2, -3}, 2, 2},
		}
		validateMaxTestCase[float32](tt, testCases)
	})
	t.Run("float64", func(tt *testing.T) {
		testCases := []testCaseMinMax[float64]{
			{make([]float64, 0), 0, -1},
			{[]float64{1}, 1, 0},
			{[]float64{0}, 0, 0},
			{[]float64{1, 0}, 1, 0},
			{[]float64{0, 1}, 1, 1},
			{[]float64{0, 1, 2, 3}, 3, 3},
			{[]float64{0, 1, 2, -3}, 2, 2},
		}
		validateMaxTestCase[float64](tt, testCases)
	})
}

func TestMin(t *testing.T) {
	t.Parallel()
	t.Run("uint", func(tt *testing.T) {
		testCases := []testCaseMinMax[uint]{
			{make([]uint, 0), 0, -1},
			{[]uint{1}, 1, 0},
			{[]uint{0}, 0, 0},
			{[]uint{1, 0}, 0, 1},
			{[]uint{0, 1}, 0, 0},
			{[]uint{3, 1, 2, 0}, 0, 3},
		}
		validateMinTestCase[uint](tt, testCases)
	})
	t.Run("int", func(tt *testing.T) {
		testCases := []testCaseMinMax[int]{
			{make([]int, 0), 0, -1},
			{[]int{1}, 1, 0},
			{[]int{0}, 0, 0},
			{[]int{1, 0}, 0, 1},
			{[]int{0, 1}, 0, 0},
			{[]int{3, 1, 2, 0}, 0, 3},
		}
		validateMinTestCase[int](tt, testCases)
	})
	t.Run("float32", func(tt *testing.T) {
		testCases := []testCaseMinMax[float32]{
			{make([]float32, 0), 0, -1},
			{[]float32{1}, 1, 0},
			{[]float32{0}, 0, 0},
			{[]float32{1, 0}, 0, 1},
			{[]float32{0, 1}, 0, 0},
			{[]float32{3, 1, 2, 0}, 0, 3},
		}
		validateMinTestCase[float32](tt, testCases)
	})
	t.Run("float64", func(tt *testing.T) {
		testCases := []testCaseMinMax[float64]{
			{make([]float64, 0), 0, -1},
			{[]float64{1}, 1, 0},
			{[]float64{0}, 0, 0},
			{[]float64{1, 0}, 0, 1},
			{[]float64{0, 1}, 0, 0},
			{[]float64{3, 1, 2, 0}, 0, 3},
		}
		validateMinTestCase[float64](tt, testCases)
	})
}
