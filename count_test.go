package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCount(t *testing.T) {
	var even = func(x int) bool {
		return x%2 == 0
	}
	assert.Equal(t, 2, Count(even)([]int{1, 2, 3, 4, 5}))
}
