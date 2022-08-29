package ramda

// Any Returns `true` if at least one of the elements of the list match the predicate, `false` otherwise.
func Any[T any](predicate func(T) bool) func([]T) bool {
	return func(list []T) bool {
		for _, it := range list {
			if predicate(it) {
				return true
			}
		}
		return false
	}
}
