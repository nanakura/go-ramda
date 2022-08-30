package ramda

import (
	"testing"
)

func TestForeach(t *testing.T) {
	ForEach[int](func(x int) {
		t.Log(x)
	})([]int{1, 2, 3})
}
