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

	fileOnly := Filter(files, func(file File) bool {
		return !file.IsDir()
	})
	l2, err := Len(fileOnly)
	if err != nil {
		t.Fatal(err)
	}
	if l2 == 0 {
		t.Fatal(err)
	}

	dirOnly := Filter(files, func(file File) bool {
		return file.IsDir()
	})
	l3, err := Len(dirOnly)
	if err != nil {
		t.Fatal(err)
	}
	if l3 == 0 {
		t.Fatal(err)
	}
	if l2+l3 != l {
		t.Fatal()
	}

}
