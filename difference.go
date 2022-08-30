package ramda

func Difference[T comparable](first []T) func([]T) []T {
	return func(second []T) []T {
		filter := map[T]bool{}
		var res []T
		for _, x := range second {
			filter[x] = true
		}
		for _, x := range first {
			if !filter[x] {
				res = append(res, x)
			}
		}
		return res
	}
}
