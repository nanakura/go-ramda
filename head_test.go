package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {
	assert.Equal(t, "fi", Head([]string{"fi", "fo", "fum"}))
}
