package ramda

import "strings"

func Trim(cutset string) func(string) string {
	return func(s string) string {
		return strings.Trim(s, cutset)
	}
}
