package ramda

import "regexp"

func Match(reg string) func(string) []string {
	return func(s string) []string {
		return regexp.MustCompile(reg).FindStringSubmatch(s)
	}
}
