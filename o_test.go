package ramda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestO(t *testing.T) {
	assert.Equal(t, 60, O[int](Multiply[int](10))(Add[int](10))(-4))
}
