package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathMod(t *testing.T) {
	assert.Equal(t, 3, MathMod(-17)(5))
	assert.Equal(t, 2, MathMod(17)(5))
}
