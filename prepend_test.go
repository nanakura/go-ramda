package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepend(t *testing.T) {
	assert.Equal(t, []string{"fee", "fi", "fo", "fum"}, Prepend("fee")([]string{"fi", "fo", "fum"}))
}
