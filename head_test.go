package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHead(t *testing.T) {
	assert.Equal(t, "fi", Head([]string{"fi", "fo", "fum"}))
}
