package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMathMod(t *testing.T) {
	assert.Equal(t, 3, MathMod(-17)(5))
	assert.Equal(t, 2, MathMod(17)(5))
}
