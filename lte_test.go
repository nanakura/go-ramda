package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLte(t *testing.T) {
	assert.Equal(t, true, Lte(1)(2))
	assert.Equal(t, false, Lte(2)(1))
}
