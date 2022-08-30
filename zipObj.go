package ramda

func ZipObj[T comparable, R any](keys []T) func([]R) map[T]R {
	return func(values []R) map[T]R {
		out := map[T]R{}
		lens := Min(len(keys))(len(values))
		for i := 0; i < lens; i++ {
			out[keys[i]] = values[i]
		}
		return out
	}
}
