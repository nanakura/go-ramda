package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	assert.Equal(t, Slice[rune](1)(3)([]rune{'a', 'b', 'c', 'd'}), []rune{'b', 'c'})
	assert.Equal(t, Slice[rune](0)(-1)([]rune{'a', 'b', 'c', 'd'}), []rune{'a', 'b', 'c'})
	assert.Equal(t, Slice[rune](-3)(-1)([]rune{'a', 'b', 'c', 'd'}), []rune{'b', 'c'})
}
