package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	assert.Equal(t, []int{1, 2}, Uniq([]int{1, 1, 2, 1}))
}

func TestUniqBy(t *testing.T) {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	assert.Equal(t, []int{-1, -5, 2, 10}, UniqBy(abs)([]int{-1, -5, 2, 10, 1, 2}))
}
