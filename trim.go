package ramda

import "strings"

func Trim() func(string) string {
	return func(s string) string {
		return strings.Trim(s, " ")
	}
}
