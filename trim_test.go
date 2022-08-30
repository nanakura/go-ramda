package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrim(t *testing.T) {
	assert.Equal(t, "xyz", Trim("   xyz    "))
}
