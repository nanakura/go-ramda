package ramda

// Curry2 Curry for 2 Params (for currying general fp functions simply)
func Curry2[T any, R any, A any, B any](fn func(A, B, ...T) R, a A, b B) func(...T) R {
	return func(args ...T) R {
		return fn(a, b, args...)
	}
}

// Curry3 Curry for 3 Params (for currying general fp functions simply)
func Curry3[T any, R any, A any, B any, C any](fn func(A, B, C, ...T) R, a A, b B, c C) func(...T) R {
	return func(args ...T) R {
		return fn(a, b, c, args...)
	}
}

// Curry4 Curry for 4 Params (for currying general fp functions simply)
func Curry4[T any, R any, A any, B any, C any, D any](fn func(A, B, C, D, ...T) R, a A, b B, c C, d D) func(...T) R {
	return func(args ...T) R {
		return fn(a, b, c, d, args...)
	}
}

// Curry5 Curry for 5 Params (for currying general fp functions simply)
func Curry5[T any, R any, A any, B any, C any, D any, E any](fn func(A, B, C, D, E, ...T) R, a A, b B, c C, d D, e E) func(...T) R {
	return func(args ...T) R {
		return fn(a, b, c, d, e, args...)
	}
}

// Curry6 Curry for 6 Params (for currying general fp functions simply)
func Curry6[T any, R any, A any, B any, C any, D any, E any, F any](fn func(A, B, C, D, E, F, ...T) R, a A, b B, c C, d D, e E, f F) func(...T) R {
	return func(args ...T) R {
		return fn(a, b, c, d, e, f, args...)
	}
}
