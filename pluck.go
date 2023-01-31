package ramda

func Pluck[V any](p int) func([][]V) []V {
	return func(list [][]V) []V {
		res := make([]V, 0, len(list))
		for _, xs := range list {
			res = Append(xs[p])(res)
		}
		return res
	}
}

func PluckObj[K comparable, V any](p K) func(map[K]map[K]V) map[K]V {
	return func(obj map[K]map[K]V) map[K]V {
		res := make(map[K]V, len(obj))
		for k, v := range obj {
			res[k] = v[p]
		}
		return res
	}
}
