package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPickBy(t *testing.T) {
	isUpperCase := func(val int, key string) bool {
		return ToUpper(key) == key
	}
	assert.Equal(t, map[string]int{"A": 3, "B": 4}, PickBy[string, int](isUpperCase)(map[string]int{"a": 1, "b": 2, "A": 3, "B": 4}))
}
