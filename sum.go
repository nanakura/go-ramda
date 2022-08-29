package ramda

func Sum[T Numeric](list []T) T {
	var res T
	for _, it := range list {
		res += it
	}
	return res
}
