package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProp(t *testing.T) {
	assert.Equal(t, 100, Prop[string, int]("x")(map[string]int{"x": 100}))
}
