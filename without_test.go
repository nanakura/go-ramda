package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithout(t *testing.T) {
	assert.Equal(t, []int{3, 4}, Without([]int{1, 2})([]int{1, 2, 1, 3, 4}))
}
