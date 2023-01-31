package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	assert.Equal(t, []string{"a"}, Repeat("a")(1))
	assert.Equal(t, []string{}, Repeat("a")(0))
	assert.Equal(t, []string{"a", "a"}, Repeat("a")(2))
}
