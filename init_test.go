package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	assert.Equal(t, Init([]int{1, 2, 3, 4}), []int{1, 2, 3})
	assert.Equal(t, Init([]int{1}), []int{})
}
