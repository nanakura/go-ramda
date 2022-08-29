package ramda

// All Dispatches to the `all` method of the second argument, if present.
func All[T any](predicate func(x T) bool) func(list []T) bool {
	return func(list []T) bool {
		for _, it := range list {
			if !predicate(it) {
				return false
			}
		}
		return true
	}
}
