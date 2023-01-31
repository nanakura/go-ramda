package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	res := Add(1)(2)
	assert.Equal(t, 3, res)
}
