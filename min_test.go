package ramda

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 123, Min(789)(123))
	assert.Equal(t, 'a', Min('a')('b'))
}

func TestMinBy(t *testing.T) {
	square := func(n int) int {
		return n * n
	}
	assert.Equal(t, 2, MinBy(square)(-3)(2))
	assert.Equal(t, 1, Reduce(func(a int, e int) int {
		return MinBy(square)(a)(e)
	})(math.MaxInt64)([]int{3, -5, 4, 1, -2}))
	assert.Equal(t, math.MaxInt64, Reduce(func(a int, e int) int {
		return MinBy(square)(a)(e)
	})(math.MaxInt64)([]int{}))
}

func TestMinInList(t *testing.T) {
	assert.Equal(t, -1, MinInList([]int{-1, 2, 3, 10}))
}
