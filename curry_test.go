package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurry2(t *testing.T) {
	add := Curry2(func(a, b int) int {
		return a + b
	})
	assert.Equal(t, 3, add(1)(2))
}
