package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNegate(t *testing.T) {
	assert.Equal(t, -42, Negate(42))
}
