package ramda

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPipe(t *testing.T) {
	var abs = func(a int) int {
		return int(math.Abs(float64(a)))
	}
	assert.Equal(t, 7, Pipe3(Multiply(2), Add(1), abs)(-4))
}
