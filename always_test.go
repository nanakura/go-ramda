package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlways(t *testing.T) {
	assert.Equal(t, "Tee", Always("Tee")())
}
