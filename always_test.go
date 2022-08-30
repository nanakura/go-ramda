package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlways(t *testing.T) {
	assert.Equal(t, "Tee", Always("Tee")())
}
