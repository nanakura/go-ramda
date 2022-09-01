package ramda

func MaxInList[T Ordered](xs []T) T {
	return Reduce[T, T](func(p, c T) T {
		return Max(p)(c)
	})(xs[0])(xs)
}
