package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, 2, Divide(4)(2))
}
