package ramda

func Reverse[T any](list []T) []T {
	length := len(list)
	for i := 0; i < length/2; i++ {
		list[i], list[length-i] = list[length-i], list[i]
	}
	return list
}
