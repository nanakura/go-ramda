package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestO(t *testing.T) {
	assert.Equal(t, 60, O[int](Multiply[int](10))(Add[int](10))(-4))
}
