package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTail(t *testing.T) {
	assert.Equal(t, []int{2, 3}, Tail([]int{1, 2, 3}))
	assert.Equal(t, []int{2}, Tail([]int{1, 2}))
	assert.Equal(t, []int{}, Tail([]int{1}))
	assert.Equal(t, []int{}, Tail([]int{}))
}
