package ramda

import "reflect"

func IsNil(o any) bool {
	if v := reflect.ValueOf(o); Kind(o) == reflect.Ptr {
		return v.IsNil()
	} else {
		return !v.IsValid()
	}
}
