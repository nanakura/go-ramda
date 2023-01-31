package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLower(t *testing.T) {
	assert.Equal(t, "abc", ToLower("ABC"))
}
