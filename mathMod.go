package ramda

func MathMod(m int) func(int) int {
	return func(p int) int {
		return ((m % p) + p) % p
	}
}
