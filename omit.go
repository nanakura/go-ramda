package ramda

import mapset "github.com/deckarep/golang-set/v2"

func Omit[K comparable, V any](names []K) func(map[K]V) map[K]V {
	return func(obj map[K]V) map[K]V {
		objHasKey := mapset.NewSet[K]()
		namesHasKey := mapset.NewSet[K]()
		res := map[K]V{}
		for k := range obj {
			objHasKey.Add(k)
		}
		for _, name := range names {
			namesHasKey.Add(name)
		}
		for k, v := range obj {
			if objHasKey.Contains(k) && !namesHasKey.Contains(k) {
				res[k] = v
			}
		}
		return res
	}
}
