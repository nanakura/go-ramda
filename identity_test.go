package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentity(t *testing.T) {
	assert.Equal(t, 1, Identity(1))
}
