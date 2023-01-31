package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimes(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 4}, Times(Identity[int])(5))
}
