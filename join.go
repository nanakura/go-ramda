package ramda

import (
	"fmt"
	"strings"
)

func Join[T any](s string) func([]T) string {
	return func(xs []T) string {
		xsToString := Map[T, string](func(it T) string {
			return fmt.Sprintf("%v", it)
		})(xs)
		return strings.Join(xsToString, s)
	}
}
