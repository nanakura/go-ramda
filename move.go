package ramda

func Move[T any](from int) func(int) func([]T) []T {
	return func(to int) func([]T) []T {
		return func(list []T) []T {
			length := len(list)
			var positiveFrom, positiveTo int
			if from < 0 {
				positiveFrom = length + from
			} else {
				positiveFrom = from
			}
			if to < 0 {
				positiveTo = length + to
			} else {
				positiveTo = to
			}
			item := list[positiveFrom]
			var result []T
			for idx := 0; idx < length; idx++ {
				if idx != positiveFrom {
					result = append(result, list[idx])
				}
			}
			if positiveFrom < 0 || positiveFrom >= length || positiveTo < 0 || positiveTo >= length {
				return list
			} else {
				return Compose3(Append[T](result[positiveTo:]...), Append(item), Append(result[:positiveTo]...))([]T{})
			}
		}
	}
}
