package ramda

// Range get an int array in given range
func Range(from int) func(int) []int {
	return func(to int) []int {
		if to >= from {
			lens := to - from
			list := make([]int, 0, lens)

			for i := from; i <= to; i++ {
				list = append(list, i)
			}
			return list
		} else {
			lens := from - to
			list := make([]int, 0, lens)

			for i := from; i >= to; i-- {
				list = append(list, i)
			}
			return list
		}
	}
}
