package ramda

func Union[T comparable](list1 []T) func([]T) []T {
	return func(list2 []T) []T {
		resMap := make(map[T]interface{})
		for _, v := range list1 {
			resMap[v] = true
		}
		for _, v := range list2 {
			resMap[v] = true
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
