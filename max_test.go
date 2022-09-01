package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 789, Max(789)(123))
	assert.Equal(t, 'b', Max('a')('b'))
}

func TestMaxBy(t *testing.T) {
	square := func(n int) int {
		return n * n
	}
	assert.Equal(t, -3, MaxBy(square)(-3)(2))
	assert.Equal(t, -5, Reduce[int, int](func(a int, e int) int {
		return MaxBy(square)(a)(e)
	})(0)([]int{3, -5, 4, 1, -2}))
	assert.Equal(t, 0, Reduce[int, int](func(a int, e int) int {
		return MaxBy(square)(a)(e)
	})(0)([]int{}))
}

func TestMaxInList(t *testing.T) {
	assert.Equal(t, 10, MaxInList([]int{1, 2, 3, 10}))
}
