package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOmit(t *testing.T) {
	assert.Equal(t, map[string]int{"b": 2, "c": 3}, Omit[string, int]([]string{"a", "d"})(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}))
}
