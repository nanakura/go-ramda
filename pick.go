package ramda

import mapset "github.com/deckarep/golang-set/v2"

func Pick[K comparable, V any](names []K) func(map[K]V) map[K]V {
	return func(obj map[K]V) map[K]V {
		namesHasKey := mapset.NewSet[K]()
		objHasKey := mapset.NewSet[K]()
		res := map[K]V{}
		for _, x := range names {
			namesHasKey.Add(x)
		}
		for k := range obj {
			objHasKey.Add(k)
		}
		for k, v := range obj {
			if namesHasKey.Contains(k) && objHasKey.Contains(k) {
				res[k] = v
			}
		}
		return res
	}
}
