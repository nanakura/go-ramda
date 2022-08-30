package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, Flatten([][]int{
		{1}, {2},
		{3, 4},
		{5},
		{6},
	}))
}
