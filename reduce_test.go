package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	assert.Equal(t, Reduce(func(acc []int, elem int) []int {
		return Append(elem)(acc)
	})([]int{})([]int{1, 2, 3}), []int{1, 2, 3})
}
