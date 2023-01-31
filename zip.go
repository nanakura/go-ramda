package ramda

func Zip[T any](l1 []T) func([]T) [][]T {
	return func(l2 []T) [][]T {
		length1 := len(l1)
		length2 := len(l2)
		if length1 == 0 || length2 == 0 {
			return [][]T{}
		}
		minLen := Min(length1)(length2)
		res := make([][]T, minLen)
		for i := 0; i < minLen; i++ {
			res[i] = []T{l1[i], l2[i]}
		}
		return res
	}
}
