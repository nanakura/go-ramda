package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMove(t *testing.T) {
	assert.Equal(t, []rune{'b', 'c', 'a', 'd', 'e', 'f'}, Move[rune](0)(2)([]rune{'a', 'b', 'c', 'd', 'e', 'f'}))
	assert.Equal(t, []rune{'f', 'a', 'b', 'c', 'd', 'e'}, Move[rune](-1)(0)([]rune{'a', 'b', 'c', 'd', 'e', 'f'}))
}
