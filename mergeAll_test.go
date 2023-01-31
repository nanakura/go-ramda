package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAll(t *testing.T) {
	assert.Equal(t, map[string]int{"foo": 1, "bar": 2, "baz": 3}, MergeAll([]map[string]int{{"foo": 1}, {"bar": 2}, {"baz": 3}}))
	assert.Equal(t, map[string]int{"foo": 2, "bar": 2}, MergeAll([]map[string]int{{"foo": 1}, {"foo": 2}, {"bar": 2}}))
}
