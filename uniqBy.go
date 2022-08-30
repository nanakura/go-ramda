package ramda

func UniqBy[T1 any, T2 comparable](fn func(T1) T2) func([]T1) []T1 {
	return func(list []T1) []T1 {
		set := map[T2]bool{}
		var result []T1
		length := len(list)
		for idx := 0; idx < length; idx++ {
			item := list[idx]
			appliedItem := fn(item)
			if !set[appliedItem] {
				set[appliedItem] = true
				result = append(result, item)
			}
		}
		return result
	}
}
