package ramda

import "strings"

func Split(s string) func(string) []string {
	return func(sep string) []string {
		return strings.Split(sep, s)
	}
}
