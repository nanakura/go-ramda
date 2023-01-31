package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, []int{4, 3, 2, 1}, Reverse([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{5, 4, 3, 2, 1}, Reverse([]int{1, 2, 3, 4, 5}))
}
