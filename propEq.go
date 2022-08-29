package ramda

func PropEq[T comparable, R comparable](val R) func(T) func(map[T]R) bool {
	return func(name T) func(map[T]R) bool {
		return func(obj map[T]R) bool {
			return Equals(val)(Prop[T, R](name)(obj))
		}
	}
}
