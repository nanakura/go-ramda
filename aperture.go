package ramda

func Aperture[T any](size int) func([]T) [][]T {
	return func(xs []T) [][]T {
		length := len(xs)
		result := make([][]T, 0, length-size+1)
		for i := 0; i < length-size+1; i++ {
			tmp := make([]T, 0, size-i)
			for j := i; j < size; j++ {
				tmp = append(tmp, xs[j])
			}
			result = append(result, tmp)
		}
		return result
	}
}
