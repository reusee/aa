package aa

import (
	"math"
	"sort"
)

func Len(a Array) (l int, err error) {
	l = sort.Search(
		int(math.MaxInt64),
		func(i int) bool {
			e := a.Get(i, nil)
			if e == ErrOutOfRange {
				return true
			} else if e != nil {
				err = e
				l = -1
				return true
			}
			return false
		},
	)
	return
}
