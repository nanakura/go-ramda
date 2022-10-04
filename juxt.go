package ramda

func Juxt[R any](fns []func(...any) R) func(...any) []R {
	return func(args ...any) []R {
		var result []R
		for _, fn := range fns {
			result = append(result, fn(args))
		}
		return result
	}
}
