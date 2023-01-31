package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLte(t *testing.T) {
	assert.Equal(t, true, Lte(1)(2))
	assert.Equal(t, false, Lte(2)(1))
}
