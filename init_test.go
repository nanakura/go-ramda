package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert.Equal(t, Init([]int{1, 2, 3, 4}), []int{1, 2, 3})
	assert.Equal(t, Init([]int{1}), []int{})
}
