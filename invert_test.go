package ramda

import (
	"testing"
)

func TestInvert(t *testing.T) {
	receResultByFirstName := map[string]string{
		"first":  "alice",
		"second": "jake",
		"third":  "alice",
	}
	t.Log(Invert(receResultByFirstName))
}
