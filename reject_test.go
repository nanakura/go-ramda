package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReject(t *testing.T) {
	isOdd := func(n int) bool {
		return n%2 != 0
	}
	assert.Equal(t, []int{2, 4}, Reject(isOdd)([]int{1, 2, 3, 4}))
}
