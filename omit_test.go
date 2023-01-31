package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOmit(t *testing.T) {
	assert.Equal(t, map[string]int{"b": 2, "c": 3}, Omit[string, int]([]string{"a", "d"})(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}))
}
