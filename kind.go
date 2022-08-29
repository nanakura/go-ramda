package ramda

import "reflect"

func Kind(obj interface{}) reflect.Kind {
	return reflect.ValueOf(obj).Kind()
}
