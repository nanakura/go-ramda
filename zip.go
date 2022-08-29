package ramda

func Zip[T comparable, R any](l1 []T) func([]R) map[T]R {
	return func(l2 []R) map[T]R {
		resMap := make(map[T]R)
		length1 := len(l1)
		length2 := len(l2)
		if length1 == 0 || length2 == 0 {
			return resMap
		}
		minLen := Min(length1)(length2)
		for i := 0; i < minLen; i++ {
			resMap[l1[i]] = l2[i]
		}
		return resMap
	}
}
