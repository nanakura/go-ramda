package ramda

func None[T any](predicate func(T) bool) func([]T) bool {
	return func(list []T) bool {
		for _, it := range list {
			if predicate(it) {
				return false
			}
		}
		return true
	}
}
