package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, Sum([]int{1, 2, 3}), 6)
}
