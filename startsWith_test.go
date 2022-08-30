package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartsWith(t *testing.T) {
	assert.Equal(t, true, StartsWith([]string{"a"})([]string{"a", "b", "c"}))
	assert.Equal(t, false, StartsWith([]string{"b"})([]string{"a", "b", "c"}))
}
