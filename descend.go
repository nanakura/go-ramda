package ramda

func Descend[T any, R Ordered](fn func(T) R) func(T) func(T) int {
	return func(a T) func(T) int {
		return func(b T) int {
			aa := fn(a)
			bb := fn(b)
			if aa > bb {
				return -1
			} else if aa < bb {
				return 1
			} else {
				return 0
			}
		}
	}
}
