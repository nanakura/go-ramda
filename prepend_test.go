package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepend(t *testing.T) {
	assert.Equal(t, []string{"fee", "fi", "fo", "fum"}, Prepend("fee")([]string{"fi", "fo", "fum"}))
}
