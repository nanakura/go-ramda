package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
	assert.Equal(t, []int{1, 2}, Difference([]int{1, 2, 3, 4})([]int{7, 6, 5, 4, 3}))
}
