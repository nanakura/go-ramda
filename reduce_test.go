package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduce(t *testing.T) {
	assert.Equal(t, Reduce(func(acc []int, elem int) []int {
		return Append(elem)(acc)
	})([]int{})([]int{1, 2, 3}), []int{1, 2, 3})
}
