package ramda

func Union[T comparable](arrList [][]T) []T {
	resultMap := make(map[T]interface{})
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	result := make([]T, len(resultMap))
	i := 0
	for k := range resultMap {
		result[i] = k
		i++
	}
	return result
}
