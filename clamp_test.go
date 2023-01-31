package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClamp(t *testing.T) {
	assert.Equal(t, 1, Clamp(1)(10)(-5))
	assert.Equal(t, 10, Clamp(1)(10)(15))
	assert.Equal(t, 4, Clamp(1)(10)(4))
}
