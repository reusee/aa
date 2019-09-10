package aa

import "reflect"

func Filter(upstream Array, fn any) *filter {
	return &filter{
		fn:       reflect.ValueOf(fn),
		upstream: upstream,
		argType:  reflect.TypeOf(fn).In(0),
	}
}

type filter struct {
	fn       reflect.Value
	upstream Array
	nextI    int
	values   []reflect.Value
	argType  reflect.Type
}

var _ Array = new(filter)

func (f *filter) Get(i int, target any) error {
start:
	if i < len(f.values) {
		if target != nil {
			reflect.ValueOf(target).Elem().Set(f.values[i])
		}
		return nil
	}

	for {
		value := reflect.New(f.argType)
		err := f.upstream.Get(f.nextI, value.Interface())
		if err == ErrOutOfRange {
			break
		}
		f.nextI++
		ok := f.fn.Call([]reflect.Value{value.Elem()})[0].Bool()
		if !ok {
			continue
		}
		f.values = append(f.values, value.Elem())
		if i < len(f.values) {
			goto start
		}
	}

	if i >= len(f.values) {
		return ErrOutOfRange
	}
	if target != nil {
		reflect.ValueOf(target).Elem().Set(f.values[i])
	}

	return nil
}

func (f *filter) Set(i int, value any) error {
	return ErrNotSupported
}
