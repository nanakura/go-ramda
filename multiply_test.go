package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	double := Multiply(2)
	triple := Multiply(3)
	assert.Equal(t, 6, double(3))
	assert.Equal(t, 12, triple(4))
	assert.Equal(t, 10, Multiply(2)(5))
}
