package ramda

import "strings"

func Split(sep string) func(string) []string {
	return func(s string) []string {
		return strings.Split(s, sep)
	}
}
