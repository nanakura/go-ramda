package ramda

func Clamp[T Numeric](min T) func(T) func(T) T {
	return func(max T) func(T) T {
		return func(value T) T {
			if value < min {
				return min
			} else if value > max {
				return max
			} else {
				return value
			}
		}
	}
}
