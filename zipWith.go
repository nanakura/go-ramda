package ramda

import (
	"math"
)

// ZipWith Creates a new list out of the two supplied by applying the function to each equally-positioned pair in the lists.
func ZipWith[A any, B any, R any](fn func(A, B) R) func([]A) func([]B) []R {
	return func(a []A) func([]B) []R {
		return func(b []B) []R {
			length := int(math.Min(float64(len(a)), float64(len(b))))
			idx := 0
			var rv []R
			for idx < length {
				rv = append(rv, fn(a[idx], b[idx]))
				idx += 1
			}
			return rv
		}
	}
}
