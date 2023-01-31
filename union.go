package ramda

func Union[T comparable](list1 []T) func([]T) []T {
	return func(list2 []T) []T {
		resMap := make(map[T]struct{})
		for _, v := range list1 {
			resMap[v] = struct{}{}
		}
		for _, v := range list2 {
			resMap[v] = struct{}{}
		}
		res := make([]T, 0, len(resMap))
		for k := range resMap {
			res = append(res, k)
		}
		return res
	}
}
