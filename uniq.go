package ramda

func Uniq[T comparable](xs []T) []T {
	return UniqBy(Identity[T])(xs)
}
