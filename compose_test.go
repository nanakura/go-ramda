package ramda

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompose(t *testing.T) {
	abs := func(a int) int {
		return int(math.Abs(float64(a)))
	}
	assert.Equal(t, 7, Compose3(abs, Add(1), Multiply(2))(-4))
}
