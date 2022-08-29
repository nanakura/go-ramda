package ramda

func IsEqual[T comparable](list1, list2 []T) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

func IsEqualMap[T comparable, R comparable](map1, map2 map[T]R) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}
