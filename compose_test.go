package ramda

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCompose(t *testing.T) {
	var abs = func(a int) int {
		return int(math.Abs(float64(a)))
	}
	assert.Equal(t, 7, Compose3(abs, Add(1), Multiply(2))(-4))
}
