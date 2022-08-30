package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	var lessThan0 = Flip(Lt[int])(0)
	var lessThan2 = Flip(Lt[int])(2)
	assert.Equal(t, false, Any(lessThan0)([]int{1, 2}))
	assert.Equal(t, true, Any(lessThan2)([]int{1, 2}))
}
