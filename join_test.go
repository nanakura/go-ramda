package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	assert.Equal(t, "1 2 3 4", Join[int](" ")([]int{1, 2, 3, 4}))
}
