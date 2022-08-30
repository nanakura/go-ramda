package ramda

func StartsWith[T comparable](prefix []T) func([]T) bool {
	return func(list []T) bool {
		lensPrefix := len(prefix)
		prefixList := Take[T](lensPrefix)(list)
		lensPrefixList := len(prefixList)
		if lensPrefix != lensPrefixList {
			return false
		}
		for idx := 0; idx < lensPrefix; idx++ {
			if prefix[idx] != prefixList[idx] {
				return false
			}
		}
		return true
	}
}
