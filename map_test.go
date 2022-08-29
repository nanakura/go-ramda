package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	assert.Equal(t, Map[int, int](func(a int) int {
		return a + 1
	})([]int{1, 2, 3}), []int{2, 3, 4})
}
