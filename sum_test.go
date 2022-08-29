package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	assert.Equal(t, Sum([]int{1, 2, 3}), 6)
}
