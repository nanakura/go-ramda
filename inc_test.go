package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInc(t *testing.T) {
	assert.Equal(t, 3, Inc(2))
}
