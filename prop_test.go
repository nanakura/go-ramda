package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProp(t *testing.T) {
	assert.Equal(t, 100, Prop[string, int]("x")(map[string]int{"x": 100}))
}
