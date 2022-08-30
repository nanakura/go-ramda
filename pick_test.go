package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPick(t *testing.T) {
	assert.Equal(t, map[string]int{"a": 1, "d": 4}, Pick[string, int]([]string{"a", "d"})(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}))
	assert.Equal(t, map[string]int{"a": 1}, Pick[string, int]([]string{"a", "e", "f"})(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}))
}
