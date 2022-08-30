package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToUpper(t *testing.T) {
	assert.Equal(t, "ABC", ToUpper("abc"))
}
