package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	even := func(x int) bool {
		return x%2 == 0
	}
	assert.Equal(t, 2, Count(even)([]int{1, 2, 3, 4, 5}))
}
