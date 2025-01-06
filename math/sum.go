package math

// Sum sums all values of artuments provided
func Sum[T Number](input ...T) (ret T) {
	for i := 0; i < len(input); i++ {
		ret = ret + input[i]
	}
	return ret
}
