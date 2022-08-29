package ramda

// Ascend Makes an asending comparator out of a function that returns a value that can be compared with '<' and '>'
func Ascend[T Ordered](fn func(any) T) func(any) func(any) int {
	return func(a any) func(any) int {
		return func(b any) int {
			aa := fn(a)
			bb := fn(b)
			if aa < bb {
				return -1
			} else if aa > bb {
				return 1
			} else {
				return 0
			}
		}
	}
}
