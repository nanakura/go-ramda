package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnion(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4}, Union([]int{1, 2, 3})([]int{2, 3, 4}))
}
