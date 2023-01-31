package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhere(t *testing.T) {
	pred := Where(map[string]func(string) bool{
		"a": Equals("foo"),
	})
	assert.Equal(t, true, pred(map[string]string{"a": "foo"}))
	assert.Equal(t, false, pred(map[string]string{"a": "xxx"}))
}
