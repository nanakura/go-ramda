package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLt(t *testing.T) {
	assert.Equal(t, true, Lt(1)(3))
	assert.Equal(t, false, Lt(3)(1))
}
