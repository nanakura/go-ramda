package ramda

// Reduce Reduce the values from left to right(func(memo,val), starting value, slice)
func Reduce[T any, R any](fn func(R, T) R) func(R) func([]T) R {
	return func(memo R) func([]T) R {
		return func(input []T) R {
			length := len(input)
			for i := 0; i < length; i++ {
				memo = fn(memo, input[i])
			}
			return memo
		}
	}
}

// ReduceIndexed Reduce the values from left to right(func(memo,val,index), starting value, slice)
func ReduceIndexed[T any, R any](fn func(R, T, int) R) func(R) func([]T) R {
	return func(memo R) func([]T) R {
		return func(input []T) R {
			length := len(input)
			for i := 0; i < length; i++ {
				memo = fn(memo, input[i], i)
			}
			return memo
		}
	}
}
