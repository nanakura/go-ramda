package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInc(t *testing.T) {
	assert.Equal(t, 3, Inc(2))
}
