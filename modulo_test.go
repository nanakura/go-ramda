package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModulo(t *testing.T) {
	assert.Equal(t, 2, Modulo(17)(3))
	assert.Equal(t, -2, Modulo(-17)(3))
	assert.Equal(t, 2, Modulo(17)(-3))
}
