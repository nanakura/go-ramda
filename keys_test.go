package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, Keys(map[string]int{"a": 1, "b": 2, "c": 3}))
}
