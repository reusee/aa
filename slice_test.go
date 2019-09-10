package aa

import "testing"

func TestSlice(t *testing.T) {
	slice := &Slice{make([]int, 0)}
	testIntArray(slice, t)
}
