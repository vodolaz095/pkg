package math

// Mean returns mean value of all elements in input slice
func Mean[T Number](input ...T) (ret T) {
	if len(input) == 0 {
		return 0
	}
	return Sum[T](input...) / T(len(input))
}
