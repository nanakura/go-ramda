package ramda

func ReduceRight[T any, R any](fn func(R, T) R) func(R) func([]T) R {
	return func(memo R) func([]T) R {
		return func(input []T) R {
			length := len(input)
			for i := length - 1; i >= 0; i-- {
				memo = fn(memo, input[i])
			}
			return memo
		}
	}
}

func ReduceRightWithIndex[T any, R any](fn func(R, T, int) R) func(R) func([]T) R {
	return func(memo R) func([]T) R {
		return func(input []T) R {
			length := len(input)
			for i := length - 1; i >= 0; i-- {
				memo = fn(memo, input[i], length-i+1)
			}
			return memo
		}
	}
}
