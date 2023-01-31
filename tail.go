package ramda

func Tail[T any](list []T) []T {
	if length := len(list); length == 1 || length == 0 {
		return []T{}
	} else {
		res := make([]T, len(list)-1)
		copy(res, list[1:])
		return res
	}
}
