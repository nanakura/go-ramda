package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, Append(3)([]int{1, 2}))
}
