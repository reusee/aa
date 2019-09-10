package aa

import (
	"math/rand"
	"testing"
)

func TestSorted(t *testing.T) {
	var ints []int64
	for i := 0; i < 5; i++ {
		ints = append(ints, rand.Int63())
	}
	sorted := Sorted(Slice(ints), func(a, b int64) bool {
		return a > b
	})
	var last int64
	for i := 0; ; i++ {
		var value int64
		if err := sorted.Get(i, &value); err == ErrOutOfRange {
			break
		} else if err != nil {
			t.Fatal(err)
		}
		if i > 0 {
			if value > last {
				t.Fatal()
			}
		}
		last = value
		_ = last
	}
}
