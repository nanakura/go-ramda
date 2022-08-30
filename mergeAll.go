package ramda

func MergeAll[K comparable, V any](list []map[K]V) map[K]V {
	resMap := map[K]V{}
	for _, x := range list {
		for k, v := range x {
			resMap[k] = v
		}
	}
	return resMap
}
