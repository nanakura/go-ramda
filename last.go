package ramda

func Last[T any](list []T) T {
	return Nth[T](-1)(list)
}
