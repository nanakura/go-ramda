package ramda

// Concat Concat slices
func Concat[T any](mine []T) func(...[]T) []T {
	return func(slices ...[]T) []T {
		mineLen := len(mine)
		totalLen := mineLen

		for _, slice := range slices {
			if slice == nil {
				continue
			}

			targetLen := len(slice)
			totalLen += targetLen
		}
		newOne := make([]T, totalLen)

		for i, item := range mine {
			newOne[i] = item
		}
		totalIndex := mineLen

		for _, slice := range slices {
			if slice == nil {
				continue
			}

			target := slice
			targetLen := len(target)
			for j, item := range target {
				newOne[totalIndex+j] = item
			}
			totalIndex += targetLen
		}

		return newOne
	}
}
