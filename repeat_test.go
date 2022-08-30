package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	assert.Equal(t, []string{"a"}, Repeat("a")(1))
	assert.Equal(t, []string{}, Repeat("a")(0))
	assert.Equal(t, []string{"a", "a"}, Repeat("a")(2))
}
