package math

// Sum sums all values of artuments provided
func Sum[T Number](input ...T) (ret T) {
	for i := range input {
		ret = ret + input[i]
	}
	return ret
}
