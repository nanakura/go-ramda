package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZip(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 4},
		{2, 5},
		{3, 6},
	}, Zip[int]([]int{1, 2, 3})([]int{4, 5, 6}))
}

func TestZipObj(t *testing.T) {
	assert.Equal(t, map[string]int{
		"a": 1, "b": 2, "c": 3,
	}, ZipObj[string, int]([]string{"a", "b", "c"})([]int{1, 2, 3}))
}
