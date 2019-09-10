package aa

import "testing"

func testIntArray(array Array, t *testing.T) {
	l, err := Len(array)
	if err != nil {
		t.Fatal(err)
	}
	if l != 0 {
		t.Fatal()
	}

	if err := array.Set(0, 42); err != nil {
		t.Fatal(err)
	}
	l, err = Len(array)
	if err != nil {
		t.Fatal(err)
	}
	if l != 1 {
		t.Fatal()
	}

	if err := array.Set(42, 0); err != nil {
		t.Fatal(err)
	}
	l, err = Len(array)
	if err != nil {
		t.Fatal(err)
	}
	if l != 43 {
		t.Fatal()
	}

}
