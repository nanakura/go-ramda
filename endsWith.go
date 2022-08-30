package ramda

func EndsWith[T comparable](suffix []T) func([]T) bool {
	return func(list []T) bool {
		lensSuffix := len(suffix)
		suffixList := TakeLast[T](lensSuffix)(list)
		lensSuffixList := len(suffixList)
		if lensSuffix != lensSuffixList {
			return false
		}
		for idx := 0; idx < lensSuffix; idx++ {
			if suffix[idx] != suffixList[idx] {
				return false
			}
		}
		return true
	}
}
