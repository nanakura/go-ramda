package ramda

func Init[T any](list []T) []T {
	length := len(list)
	if length == 1 || length == 0 {
		return []T{}
	}
	return list[0 : length-1]
}
