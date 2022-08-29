package ramda

import (
	"testing"
)

func TestForeach(t *testing.T) {
	ForEach[int](func(t int, i int) {
		println(t)
	})([]int{1, 2, 3})
}
