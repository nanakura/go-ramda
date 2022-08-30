package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppend(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, Append(3)([]int{1, 2}))
}
