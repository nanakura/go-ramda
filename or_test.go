package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOr(t *testing.T) {
	assert.Equal(t, true, Or(true)(true))
	assert.Equal(t, true, Or(true)(false))
	assert.Equal(t, true, Or(false)(true))
	assert.Equal(t, false, Or(false)(false))
}
