package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartsWith(t *testing.T) {
	assert.Equal(t, true, StartsWith([]string{"a"})([]string{"a", "b", "c"}))
	assert.Equal(t, false, StartsWith([]string{"b"})([]string{"a", "b", "c"}))
}
