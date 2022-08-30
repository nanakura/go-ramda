package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNot(t *testing.T) {
	assert.Equal(t, false, Not(true))
	assert.Equal(t, true, Not(false))
}
