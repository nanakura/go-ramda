package ramda

func DropWhile[T any](f func(T) bool, list ...T) []T {
	if f == nil {
		return make([]T, 0)
	}
	var newList []T
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]T, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}
