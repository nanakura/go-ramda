package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndsWith(t *testing.T) {
	assert.Equal(t, true, EndsWith([]string{"c"})([]string{"a", "b", "c"}))
	assert.Equal(t, false, EndsWith([]string{"b"})([]string{"a", "b", "c"}))
}
