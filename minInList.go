package ramda

func MinInList[T Ordered](xs []T) T {
	return Reduce[T, T](func(p, c T) T {
		return Min(p)(c)
	})(xs[0])(xs)
}
