package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqBy(t *testing.T) {
	assert.Equal(t, true, EqBy(func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	})(-5)(5))
}
