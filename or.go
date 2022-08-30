package ramda

func Or(a bool) func(bool) bool {
	return func(b bool) bool {
		return a || b
	}
}
