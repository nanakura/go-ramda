package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	xs := []map[string]int{
		{"a": 1},
		{"a": 2},
		{"a": 3},
	}
	assert.Equal(t, map[string]int{"a": 2}, Find(PropEq[string, int](2)("a"))(xs))
}

func TestFindIndex(t *testing.T) {
	xs := []map[string]int{
		{"a": 1},
		{"a": 2},
		{"a": 3},
	}
	assert.Equal(t, 1, FindIndex(PropEq[string, int](2)("a"))(xs))
	assert.Equal(t, -1, FindIndex(PropEq[string, int](4)("a"))(xs))
}

func TestFindLast(t *testing.T) {
	xs := []map[string]int{
		{
			"a": 1,
			"b": 0,
		},
		{
			"a": 1,
			"b": 1,
		},
	}
	assert.Equal(t, map[string]int{"a": 1, "b": 1}, FindLast(PropEq[string, int](1)("a"))(xs))
	assert.Equal(t, nil, FindLast(PropEq[string, int](4)("a"))(xs))
}

func TestFindLastIndex(t *testing.T) {
	xs := []map[string]int{
		{
			"a": 1,
			"b": 0,
		},
		{
			"a": 1,
			"b": 1,
		},
	}
	assert.Equal(t, 1, FindLastIndex(PropEq[string, int](1)("a"))(xs))
	assert.Equal(t, -1, FindLastIndex(PropEq[string, int](4)("a"))(xs))
}
