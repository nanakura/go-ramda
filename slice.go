package ramda

func Slice[T any](fromIndex int) func(int) func([]T) []T {
	return func(toIndex int) func([]T) []T {
		return func(list []T) []T {
			lens := len(list)
			if fromIndex < 0 {
				fromIndex = lens + fromIndex
			}
			if toIndex < 0 {
				toIndex = lens + toIndex
			}
			if toIndex < fromIndex {
				fromIndex, toIndex = toIndex, fromIndex
			}
			return list[fromIndex:toIndex]
		}
	}
}
