package ramda

func Aperture[T any](size int) func([]T) [][]T {
	return func(xs []T) [][]T {
		var result [][]T
		length := len(xs)
		for i := 0; i < length-size+1; i++ {
			var tmp []T
			for j := i; j < size; j++ {
				tmp = append(tmp, xs[j])
			}
			result = append(result, tmp)
		}
		return result
	}
}
