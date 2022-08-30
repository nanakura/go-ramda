package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimes(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 4}, Times(Identity[int])(5))
}
