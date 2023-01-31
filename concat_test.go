package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	assert.Equal(t, []int{4, 5, 6, 1, 2, 3}, Concat([]int{4, 5, 6})([]int{1, 2, 3}))
}
