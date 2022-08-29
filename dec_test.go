package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDec(t *testing.T) {
	assert.Equal(t, Dec(42), 41)
}
