package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluck(t *testing.T) {
	assert.Equal(t, []int{1, 3}, Pluck[int](0)([][]int{{1, 2}, {3, 4}}))
}

func TestPluckObj(t *testing.T) {
	assert.Equal(t, map[string]int{"a": 3, "b": 5}, PluckObj[string, int]("val")(map[string]map[string]int{"a": {"val": 3}, "b": {"val": 5}}))
}
