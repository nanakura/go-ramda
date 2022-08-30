package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNth(t *testing.T) {
	list := []string{"foo", "bar", "baz", "quux"}
	assert.Equal(t, "bar", Nth[string](1)(list))
	assert.Equal(t, "quux", Nth[string](-1)(list))
}
