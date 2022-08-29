package ramda

func Compose[T any](fnList ...func(...T) []T) func(...T) []T {
	return func(s ...T) []T {
		f := fnList[0]
		nextFnList := fnList[1:]

		if len(fnList) == 1 {
			return f(s...)
		}

		return f(Compose(nextFnList...)(s...)...)
	}
}

func ComposeInterface(fnList ...func(...interface{}) []interface{}) func(...interface{}) []interface{} {
	return Compose(fnList...)
}
