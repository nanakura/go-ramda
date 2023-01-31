package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlip(t *testing.T) {
	assert.Equal(t, true, Flip(Lt[int])(2)(1))
}
