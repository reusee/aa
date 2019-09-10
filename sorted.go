package aa

import (
	"reflect"
	"sort"
	"sync"
)

type sorted struct {
	compare   reflect.Value
	upstream  Array
	indexes   []int
	sortOnce  sync.Once
	valueType reflect.Type
	err       error
}

func Sorted(upstream Array, compare any) *sorted {
	return &sorted{
		upstream:  upstream,
		compare:   reflect.ValueOf(compare),
		valueType: reflect.TypeOf(compare).In(0),
	}
}

var _ Array = new(sorted)

func (s *sorted) Get(i int, target any) error {
	s.sortOnce.Do(func() {
		l, err := Len(s.upstream)
		if err != nil {
			s.err = err
			return
		}
		for i := 0; i < l; i++ {
			s.indexes = append(s.indexes, i)
		}
		sort.Slice(s.indexes, func(i, j int) bool {
			a := reflect.New(s.valueType)
			if err := s.upstream.Get(s.indexes[i], a.Interface()); err != nil {
				s.err = err
				return false
			}
			b := reflect.New(s.valueType)
			if err := s.upstream.Get(s.indexes[j], b.Interface()); err != nil {
				s.err = err
				return false
			}
			return s.compare.Call([]reflect.Value{a.Elem(), b.Elem()})[0].Bool()
		})
	})

	if s.err != nil {
		return s.err
	}
	if i >= len(s.indexes) {
		return ErrOutOfRange
	}

	return s.upstream.Get(s.indexes[i], target)
}

func (s *sorted) Set(i int, value any) error {
	return s.upstream.Set(i, value)
}
