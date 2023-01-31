package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqBy(t *testing.T) {
	assert.Equal(t, true, EqBy(func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	})(-5)(5))
}
