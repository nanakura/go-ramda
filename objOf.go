package ramda

func ObjOf[K comparable, V any](key K) func(V) map[K]V {
	return func(val V) map[K]V {
		res := map[K]V{}
		res[key] = val
		return res
	}
}
