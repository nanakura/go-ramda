package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUpper(t *testing.T) {
	assert.Equal(t, "ABC", ToUpper("abc"))
}
