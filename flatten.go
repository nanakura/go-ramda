package ramda

func Flatten[T any](list [][]T) []T {
	result := make([]T, 0)

	// for _, v := range list {
	// 	result = append(result, v...)
	// }

	return Concat(result)(list...)
}
