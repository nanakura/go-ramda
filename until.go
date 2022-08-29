package ramda

func Until[T any](pred func(T) bool) func(func(T) T) func(T) T {
	return func(fn func(T) T) func(T) T {
		return func(init T) T {
			val := init
			for !pred(val) {
				val = fn(val)
			}
			return val
		}
	}
}
