package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLt(t *testing.T) {
	assert.Equal(t, true, Lt(1)(3))
	assert.Equal(t, false, Lt(3)(1))
}
