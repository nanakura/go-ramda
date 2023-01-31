package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModulo(t *testing.T) {
	assert.Equal(t, 2, Modulo(17)(3))
	assert.Equal(t, -2, Modulo(-17)(3))
	assert.Equal(t, 2, Modulo(17)(-3))
}
