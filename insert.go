package ramda

func Insert[T any](idx int) func(T) func([]T) []T {
	return func(elt T) func([]T) []T {
		return func(xs []T) []T {
			lens := len(xs)
			var res []T
			for i := 0; i < idx; i++ {
				res = append(res, xs[i])
			}
			res = append(res, elt)
			for i := idx; i < lens; i++ {
				res = append(res, xs[i])
			}
			return res
		}
	}
}
