package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNegate(t *testing.T) {
	assert.Equal(t, -42, Negate(42))
}
