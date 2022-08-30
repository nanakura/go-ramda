package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOn(t *testing.T) {
	var eqBy = On[int, bool, map[string]int](func(a, b int) bool {
		return a == b
	})
	assert.Equal(t, true, eqBy(Prop[string, int]("a"))(map[string]int{"b": 0, "a": 1})(map[string]int{"a": 1}))
}
