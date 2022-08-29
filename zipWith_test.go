package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func f(a int, b int) int {
	return a + b
}

func TestZipWith(t *testing.T) {
	assert.Equal(t, ZipWith(f)([]int{1, 2, 3})([]int{1, 2, 3}), []int{2, 4, 6})
}
