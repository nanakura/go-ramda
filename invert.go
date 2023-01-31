package ramda

func Invert[T1 comparable, T2 comparable](obj map[T1]T2) map[T2][]T1 {
	props := Keys(obj)
	lens := len(props)
	out := make(map[T2][]T1, lens)
	for idx := 0; idx < lens; idx++ {
		key := props[idx]
		val := obj[key]
		out[val] = append(out[val], key)
	}
	return out
}
