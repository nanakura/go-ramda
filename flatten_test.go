package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, Flatten([][]int{
		{1},
		{2},
		{3, 4},
		{5},
		{6},
	}))
}
