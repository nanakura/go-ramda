package ramda

import mapset "github.com/deckarep/golang-set/v2"

func Without[T comparable](list1 []T) func([]T) []T {
	return func(list2 []T) []T {
		toRemove := mapset.NewSet[T]()
		var res []T
		for _, it := range list1 {
			toRemove.Add(it)
		}
		for _, it := range list2 {
			if !toRemove.Contains(it) {
				res = append(res, it)
			}
		}
		return res
	}
}
