package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhere(t *testing.T) {
	var pred = Where(map[string]func(string) bool{
		"a": Equals("foo"),
	})
	assert.Equal(t, true, pred(map[string]string{"a": "foo"}))
	assert.Equal(t, false, pred(map[string]string{"a": "xxx"}))
}
