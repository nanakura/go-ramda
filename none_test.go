package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNone(t *testing.T) {
	isEven := func(n int) bool {
		return n%2 == 0
	}
	isOdd := func(n int) bool {
		return !isEven(n)
	}
	assert.Equal(t, true, None(isEven)([]int{1, 3, 5, 7, 9, 11}))
	assert.Equal(t, false, None(isOdd)([]int{1, 3, 5, 7, 9, 11}))
}
