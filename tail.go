package ramda

func Tail[T any](list []T) []T {
	if length := len(list); length == 1 || length == 0 {
		return []T{}
	} else {
		return list[1:]
	}
}
