package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c", "xyz", "d"}, Split(".")("a.b.c.xyz.d"))
}
