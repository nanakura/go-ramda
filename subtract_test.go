package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	assert.Equal(t, 0, Subtract(1)(1))
}
