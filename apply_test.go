package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	max := func(xs ...int) int {
		res := xs[0]
		for _, x := range xs {
			if x > res {
				res = x
			}
		}
		return res
	}
	assert.Equal(t, 42, Apply(max)([]int{1, 2, 3, -99, 42, 6, 7}))
}
