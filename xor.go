package ramda

func Xor(a bool) func(bool) bool {
	return func(b bool) bool {
		var numA int
		var numB int
		if !a {
			numA = 1
		} else {
			numB = 0
		}
		if !b {
			numB = 1
		} else {
			numB = 0
		}
		res := numA ^ numB
		if res == 1 {
			return true
		} else {
			return false
		}
	}
}
