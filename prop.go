package ramda

func Prop[T comparable, R any](key T) func(map[T]R) R {
	return func(m map[T]R) R {
		return m[key]
	}
}

// TODO PropForInterface
//func PropForInterface[R any](key string) func(any) R {
//	return func(obj any) R {
//		reflect.ValueOf(obj).FieldByName(key)
//	}
//}
