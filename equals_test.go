package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquals(t *testing.T) {
	assert.Equal(t, true, Equals(2)(2))
	assert.Equal(t, false, Equals(2)(3))
}
