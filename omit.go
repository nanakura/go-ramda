package ramda

func Omit[K comparable, V any](names []K) func(map[K]V) map[K]V {
	return func(obj map[K]V) map[K]V {
		objHasKey := map[K]bool{}
		namesHasKey := map[K]bool{}
		res := map[K]V{}
		for k := range obj {
			objHasKey[k] = true
		}
		for _, name := range names {
			namesHasKey[name] = true
		}
		for k, v := range obj {
			if objHasKey[k] && !namesHasKey[k] {
				res[k] = v
			}
		}
		return res
	}
}
