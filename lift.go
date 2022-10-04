package ramda

// TODO
func Lift[T any](fn func(...T)) func(...[]T) []T {
	return func(lists ...[]T) []T {
		// TODO length panic
		resultLength := 1
		listsLength := len(lists)
		lens := make([]int, listsLength)
		for idx, list := range lists {
			lens[idx] = len(list)
			resultLength = resultLength * lens[idx]
		}
		result := make([]T, resultLength)
		argsList := make([][]T, resultLength)
		for i := 0; i < listsLength; i++ {
			var tmp []T
			for j := 0; j < lens[i]; j++ {
				for k := 0; k < listsLength; k++ {
					// ?
					tmp = append(tmp, lists[j][i])
				}
			}
			argsList = append(argsList, tmp)
		}
		return result
	}
}
