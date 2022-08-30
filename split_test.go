package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplit(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c", "xyz", "d"}, Split(".")("a.b.c.xyz.d"))
}
