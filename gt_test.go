package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGt(t *testing.T) {
	assert.Equal(t, true, Gt(3)(0))
	assert.Equal(t, false, Gt(3)(4))
}

func TestGte(t *testing.T) {
	assert.Equal(t, true, Gte(3)(3))
	assert.Equal(t, true, Gte(3)(1))
	assert.Equal(t, false, Gte(3)(4))
}
