package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTake(t *testing.T) {
	assert.Equal(t, []string{"foo"}, Take[string](1)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{"foo", "bar"}, Take[string](2)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{"foo", "bar", "baz"}, Take[string](3)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{"foo", "bar", "baz"}, Take[string](4)([]string{"foo", "bar", "baz"}))
}

func TestTakeLast(t *testing.T) {
	assert.Equal(t, []string{"baz"}, TakeLast[string](1)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{"bar", "baz"}, TakeLast[string](2)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{"foo", "bar", "baz"}, TakeLast[string](3)([]string{"foo", "bar", "baz"}))
	assert.Equal(t, []string{"foo", "bar", "baz"}, TakeLast[string](4)([]string{"foo", "bar", "baz"}))
}
