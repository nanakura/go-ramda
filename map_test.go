package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert.Equal(t, Map[int, int](func(a int) int {
		return a + 1
	})([]int{1, 2, 3}), []int{2, 3, 4})
}
