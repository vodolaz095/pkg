package math

// Derivative returns 1st derivative in assumption input is array of values of fixed step argument function
func Derivative[T Number](input []T) (ret []T) {
	ret = make([]T, len(input))
	var acc T = 0
	for i := 0; i < len(input); i++ {
		ret[i] = input[i] - acc
		acc = input[i]
	}
	return ret
}
