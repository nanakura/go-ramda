package ramda

func On[A, B, C any](f func(A, A) B) func(func(C) A) func(C) func(C) B {
	return func(g func(C) A) func(C) func(C) B {
		return func(a C) func(C) B {
			return func(b C) B {
				return f(g(a), g(b))
			}
		}
	}
}
