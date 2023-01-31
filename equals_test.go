package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquals(t *testing.T) {
	assert.Equal(t, true, Equals(2)(2))
	assert.Equal(t, false, Equals(2)(3))
}
