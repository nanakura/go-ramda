package ramda

func UnionWith[T comparable](pred func(T) bool) func([]T) func([]T) []T {
	return func(list1 []T) func([]T) []T {
		return func(list2 []T) []T {
			resMap := make(map[T]struct{})
			for _, v := range list1 {
				if pred(v) {
					resMap[v] = struct{}{}
				}
			}
			for _, v := range list2 {
				if pred(v) {
					resMap[v] = struct{}{}
				}
			}
			res := make([]T, len(resMap))
			i := 0
			for k := range resMap {
				res[i] = k
				i++
			}
			return res
		}
	}
}
