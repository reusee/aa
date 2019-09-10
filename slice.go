package aa

import "reflect"

type slice struct {
	any
}

var _ Array = new(slice)

func Slice(s any) *slice {
	return &slice{
		any: s,
	}
}

func (s slice) Get(i int, target any) error {
	v := reflect.ValueOf(s.any)
	if i >= v.Len() {
		return ErrOutOfRange
	}
	if target != nil {
		reflect.ValueOf(target).Elem().Set(v.Index(i))
	}
	return nil
}

func (s *slice) Set(i int, value any) error {
	v := reflect.ValueOf(s.any)
	if i >= v.Len() {
		newArray := reflect.MakeSlice(v.Type(), i+1, i+1)
		reflect.Copy(newArray, v)
		v = newArray
		s.any = v.Interface()
	}
	v.Index(i).Set(reflect.ValueOf(value))
	return nil
}
