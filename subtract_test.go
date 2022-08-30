package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubtract(t *testing.T) {
	assert.Equal(t, 0, Subtract(1)(1))
}
