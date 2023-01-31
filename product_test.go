package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	assert.Equal(t, 38400, Product([]int{2, 4, 6, 8, 100, 1}))
}
