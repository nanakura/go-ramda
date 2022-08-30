package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurry2(t *testing.T) {
	var add = Curry2(func(a, b int) int {
		return a + b
	})
	assert.Equal(t, 3, add(1)(2))
}
