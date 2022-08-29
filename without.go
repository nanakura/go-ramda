package ramda

func Without[T comparable](list1 []T) func([]T) []T {
	return func(list2 []T) []T {
		toRemove := map[T]bool{}
		var res []T
		for _, it := range list1 {
			toRemove[it] = true
		}
		for _, it := range list2 {
			if !toRemove[it] {
				res = append(res, it)
			}
		}
		return res
	}
}
