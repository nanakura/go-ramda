package ramda

func When[T1 any, T2 any](pred func(T1) bool) func(func(T1) T2) func(T1) any {
	return func(whenTrueFn func(T1) T2) func(T1) any {
		return func(x T1) any {
			if pred(x) {
				return whenTrueFn(x)
			} else {
				return x
			}
		}
	}
}
