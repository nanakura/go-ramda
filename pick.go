package ramda

func Pick[K comparable, V any](names []K) func(map[K]V) map[K]V {
	return func(obj map[K]V) map[K]V {
		namesHasKey := map[K]bool{}
		objHasKey := map[K]bool{}
		res := map[K]V{}
		for _, x := range names {
			namesHasKey[x] = true
		}
		for k := range obj {
			objHasKey[k] = true
		}
		for k, v := range obj {
			if namesHasKey[k] && objHasKey[k] {
				res[k] = v
			}
		}
		return res
	}
}
