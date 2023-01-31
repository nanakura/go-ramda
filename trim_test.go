package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrim(t *testing.T) {
	assert.Equal(t, "xyz", Trim("   xyz    "))
}
