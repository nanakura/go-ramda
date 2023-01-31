package ramda

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	arr := Union([]int{1, 2, 3})([]int{2, 3, 4})
	sort.Ints(arr)
	assert.Equal(t, arr, []int{1, 2, 3, 4})
}
