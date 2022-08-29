package ramda

// TODO fix
func Range[T Numeric](from T) func(T) []T {
	return func(to T) []T {
		if to >= from {
			lens := to - from
			var list []T
			for i, idx := from, 0; i < lens; idx++ {
				list = append(list, i)
				i++
			}
			return list
		} else {
			lens := from - to
			var list []T
			for i, idx := to, 0; i < lens; idx++ {
				list = append(list, i)
				i++
			}
			return list
		}
	}
}
