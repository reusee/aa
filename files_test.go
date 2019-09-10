package aa

import (
	"testing"
)

func TestFiles(t *testing.T) {
	files, err := NewFiles(".")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; ; i++ {
		var info File
		if err := files.Get(i, &info); err == ErrOutOfRange {
			break
		} else if err != nil {
			t.Fatal(err)
		}
	}
	l, err := Len(files)
	if err != nil {
		t.Fatal(err)
	}
	if l == 0 {
		t.Fatal()
	}
}
