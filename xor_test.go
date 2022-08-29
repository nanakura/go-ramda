package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXor(t *testing.T) {
	assert.Equal(t, Xor(true)(true), false)
	assert.Equal(t, Xor(true)(false), true)
	assert.Equal(t, Xor(false)(true), true)
	assert.Equal(t, Xor(false)(false), false)
}
