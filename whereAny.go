package ramda

func WhereAny[K comparable, V any](spec map[K]func(V) bool) func(map[K]V) bool {
	return func(testObj map[K]V) bool {
		for prop := range spec {
			if spec[prop](testObj[prop]) {
				return true
			}
		}
		return false
	}
}
