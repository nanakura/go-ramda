package ramda

func Swap[T any](indexA int) func(int) func([]T) []T {
	return func(indexB int) func([]T) []T {
		return func(o []T) []T {
			length := len(o)
			if indexA < 0 {
				indexA = length + indexA
			}
			if indexB < 0 {
				indexB = length + indexB
			}
			o[indexA], o[indexB] = o[indexB], o[indexA]
			return o
		}
	}
}
