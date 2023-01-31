package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	assert.Equal(t, true, All(Equals(3))([]int{3, 3, 3, 3}))
	assert.Equal(t, false, All(Equals(3))([]int{3, 3, 1, 3}))
}
