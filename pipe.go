package ramda

// Pipe Pipe the functions from left to right
func Pipe[T any](fnList ...func(...T) []T) func(...T) []T {
	return func(s ...T) []T {
		lastIndex := len(fnList) - 1
		f := fnList[lastIndex]
		nextFnList := fnList[:lastIndex]

		if len(fnList) == 1 {
			return f(s...)
		}

		return f(Pipe(nextFnList...)(s...)...)
	}
}

// PipeInterface Pipe the functions from left to right
func PipeInterface(fnList ...func(...interface{}) []interface{}) func(...interface{}) []interface{} {
	return Pipe(fnList...)
}
