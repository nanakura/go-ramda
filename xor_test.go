package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXor(t *testing.T) {
	assert.Equal(t, Xor(true)(true), false)
	assert.Equal(t, Xor(true)(false), true)
	assert.Equal(t, Xor(false)(true), true)
	assert.Equal(t, Xor(false)(false), false)
}
