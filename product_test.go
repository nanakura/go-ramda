package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct(t *testing.T) {
	assert.Equal(t, 38400, Product([]int{2, 4, 6, 8, 100, 1}))
}
