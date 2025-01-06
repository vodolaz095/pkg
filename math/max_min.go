package math

// Max returns value and index of biggest element in input slice
func Max[T Number](input ...T) (maximum T, index int) {
	var i int
	if len(input) == 0 {
		return 0, -1
	}
	maximum = input[0]
	for i = 0; i < len(input); i++ {
		if input[i] > maximum {
			maximum = input[i]
			index = i
		}
	}
	return maximum, index
}

// Min returns value and index of smallest element in input slice
func Min[T Number](input ...T) (minimum T, index int) {
	var i int
	if len(input) == 0 {
		return 0, -1
	}
	minimum = input[0]
	for i = 0; i < len(input); i++ {
		if input[i] < minimum {
			minimum = input[i]
			index = i
		}
	}
	return minimum, index
}
