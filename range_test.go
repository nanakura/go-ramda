package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange(t *testing.T) {
	assert.Equal(t, Range(1)(3), []int{1, 2, 3})
	assert.Equal(t, Range(-3)(-1), []int{-3, -2, -1})
	assert.Equal(t, Range(-3)(3), []int{-3, -2, -1, 0, 1, 2, 3})
	assert.Equal(t, Range(3)(1), []int{3, 2, 1})
	assert.Equal(t, Range(-1)(-3), []int{-1, -2, -3})
	assert.Equal(t, Range(3)(-3), []int{3, 2, 1, 0, -1, -2, -3})
}
