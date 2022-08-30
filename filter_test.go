package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	var isEven = func(n int) bool {
		return n%2 == 0
	}
	assert.Equal(t, []int{2, 4}, Filter(isEven)([]int{1, 2, 3, 4}))
}
