package ramda

func Where[T any](spec map[string]func(T) bool) func(map[string]T) bool {
	return func(testObj map[string]T) bool {
		for prop := range spec {
			if !spec[prop](testObj[prop]) {
				return false
			}
		}
		return true
	}
}
