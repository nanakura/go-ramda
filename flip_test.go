package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlip(t *testing.T) {
	assert.Equal(t, true, Flip(Lt[int])(2)(1))
}
