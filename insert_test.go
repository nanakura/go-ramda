package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 3, 4}, Insert[int](2)(3)([]int{1, 2, 3, 4}))
}
