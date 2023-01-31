package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNot(t *testing.T) {
	assert.Equal(t, false, Not(true))
	assert.Equal(t, true, Not(false))
}
