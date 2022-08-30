package ramda

func Where[K comparable, V any](spec map[K]func(V) bool) func(map[K]V) bool {
	return func(testObj map[K]V) bool {
		for prop := range spec {
			if !spec[prop](testObj[prop]) {
				return false
			}
		}
		return true
	}
}
