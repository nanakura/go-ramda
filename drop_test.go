package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDrop(t *testing.T) {
	assert.Equal(t, []string{"bar", "baz"}, Drop[string](1)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{"baz"}, Drop[string](2)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{}, Drop[string](3)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{}, Drop[string](4)([]string{"foo", "bar", "baz"}))
}

func TestDropLast(t *testing.T) {
	assert.Equal(t, []string{"foo", "bar"}, DropLast[string](1)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{"foo"}, DropLast[string](2)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{}, DropLast[string](3)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{}, DropLast[string](4)([]string{"foo", "bar", "baz"}))
}

func TestDropWhile(t *testing.T) {
	var lteTwo = func(x int) bool {
		return x <= 2
	}
	assert.Equal(t, []int{3, 4, 3, 2, 1}, DropWhile(lteTwo)([]int{1, 2, 3, 4, 3, 2, 1}))
}
