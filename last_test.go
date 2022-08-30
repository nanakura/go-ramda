package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLast(t *testing.T) {
	assert.Equal(t, 3, Last([]int{1, 2, 3}))
}

func TestLastIndexOf(t *testing.T) {
	assert.Equal(t, 6, LastIndexOf(3)([]int{-1, 3, 3, 0, 1, 2, 3, 4}))
	assert.Equal(t, -1, LastIndexOf(100)([]int{-1, 3, 3, 0, 1, 2, 3, 4}))
}
