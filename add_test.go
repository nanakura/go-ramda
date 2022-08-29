package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	res := Add(1)(2)
	assert.Equal(t, 3, res)
}
