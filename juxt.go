package ramda

func Juxt[R any](fns []func(...any) R) func(...any) []R {
	return func(args ...any) []R {
		result := make([]R, 0, len(fns))
		for _, fn := range fns {
			result = append(result, fn(args))
		}
		return result
	}
}
