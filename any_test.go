package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	lessThan0 := Flip(Lt[int])(0)
	lessThan2 := Flip(Lt[int])(2)
	assert.Equal(t, false, Any(lessThan0)([]int{1, 2}))
	assert.Equal(t, true, Any(lessThan2)([]int{1, 2}))
}
