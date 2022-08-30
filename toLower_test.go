package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToLower(t *testing.T) {
	assert.Equal(t, "abc", ToLower("ABC"))
}
