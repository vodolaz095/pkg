package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tableFunctionTestCase[T Number] struct {
	input  []T
	output []T
}

func validateDerivative[T Number](t *testing.T, testCases []tableFunctionTestCase[T]) {
	var output []T
	for i := range testCases {
		output = Derivative[T](testCases[i].input)
		assert.Equal(t, len(testCases[i].input), len(testCases[i].output),
			"test case %v has different input and output length", i)
		assert.Equal(t, len(testCases[i].input), len(output),
			"test case %v received wrong output length", i)

		for j := range testCases[i].output {
			assert.Equal(t, testCases[i].output[j], output[j],
				"test case %v - wrong element returned on index %v", i, j,
			)
		}
	}
}

func TestDerivative(t *testing.T) {
	t.Run("int", func(tt *testing.T) {
		testCases := []tableFunctionTestCase[int]{
			{make([]int, 0), make([]int, 0)},
			{[]int{1}, []int{1}},
			{
				[]int{1, 2, 3}, // f(n) = n + 1
				[]int{1, 1, 1}, // f'(n) = 1
			},
			{
				[]int{2, 4, 6}, // f(n) = n + 2
				[]int{2, 2, 2}, // f'(n) = 1
			},
			{
				[]int{1, 4, 9, 16}, // f(n) = n^2
				[]int{1, 3, 5, 7},  // f'(n) = 2n-1
			},
			{
				Sequence[int](3, func(i int) int {
					return i + 1
				}),
				Sequence[int](3, func(i int) int {
					return 1
				}),
			},
			{
				Sequence[int](100, func(i int) int {
					return i + 1
				}),
				Sequence[int](100, func(i int) int {
					return 1
				}),
			},
			{
				Sequence[int](100, func(i int) int {
					return i * i
				}),
				Sequence[int](100, func(i int) int {
					if i == 0 {
						return 0
					}
					return 2*i - 1
				}),
			},
		}
		validateDerivative[int](tt, testCases)
	})
}
