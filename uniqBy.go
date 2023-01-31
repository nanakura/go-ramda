package ramda

import mapset "github.com/deckarep/golang-set/v2"

func UniqBy[T1 any, T2 comparable](fn func(T1) T2) func([]T1) []T1 {
	return func(list []T1) []T1 {
		set := mapset.NewSet[T2]()
		var result []T1
		length := len(list)
		for idx := 0; idx < length; idx++ {
			item := list[idx]
			appliedItem := fn(item)
			if !set.Contains(appliedItem) {
				set.Add(appliedItem)
				result = append(result, item)
			}
		}
		return result
	}
}
