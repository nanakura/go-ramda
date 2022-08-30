package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIdentity(t *testing.T) {
	assert.Equal(t, 1, Identity(1))
}
