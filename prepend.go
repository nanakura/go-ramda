package ramda

func Prepend[T any](el T) func([]T) []T {
	return func(list []T) []T {
		return append([]T{el}, list...)
	}
}
