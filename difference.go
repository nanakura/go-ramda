package ramda

import mapset "github.com/deckarep/golang-set/v2"

func Difference[T comparable](first []T) func([]T) []T {
	return func(second []T) []T {
		filter := mapset.NewSet[T]()
		var res []T
		for _, x := range second {
			filter.Add(x)
		}
		for _, x := range first {
			if !filter.Contains(x) {
				res = append(res, x)
			}
		}
		return res
	}
}
