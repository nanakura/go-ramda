package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeys(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, Keys(map[string]int{"a": 1, "b": 2, "c": 3}))
}
