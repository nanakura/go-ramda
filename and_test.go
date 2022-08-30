package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnd(t *testing.T) {
	assert.Equal(t, true, And(true)(true))
	assert.Equal(t, false, And(true)(false))
	assert.Equal(t, false, And(false)(true))
	assert.Equal(t, false, And(false)(false))
}
