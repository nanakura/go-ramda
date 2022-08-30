package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoin(t *testing.T) {
	assert.Equal(t, "1 2 3 4", Join[int](" ")([]int{1, 2, 3, 4}))
}
