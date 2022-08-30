package ramda

func PickBy[K comparable, V any](test func(V, K) bool) func(map[K]V) map[K]V {
	return func(obj map[K]V) map[K]V {
		res := map[K]V{}
		for prop := range obj {
			if test(obj[prop], prop) {
				res[prop] = obj[prop]
			}
		}
		return res
	}
}
