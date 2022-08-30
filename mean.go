package ramda

func Mean(list []int) int {
	return Sum(list) / len(list)
}
