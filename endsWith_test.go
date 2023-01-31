package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndsWith(t *testing.T) {
	assert.Equal(t, true, EndsWith([]string{"c"})([]string{"a", "b", "c"}))
	assert.Equal(t, false, EndsWith([]string{"b"})([]string{"a", "b", "c"}))
}
