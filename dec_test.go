package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec(t *testing.T) {
	assert.Equal(t, Dec(42), 41)
}
