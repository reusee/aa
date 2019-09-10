package aa

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
)

type Files struct {
	queue      []*os.File
	queuePaths []string
	files      []File
	err        error
}

type File struct {
	Dir string
	os.FileInfo
}

func NewFiles(dir string) (*Files, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	abs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	return &Files{
		queue:      []*os.File{f},
		queuePaths: []string{abs},
	}, nil
}

func (f *Files) Get(i int, target any) error {
	if f.err != nil {
		return f.err
	}

	if i < len(f.files) {
		if target != nil {
			reflect.ValueOf(target).Elem().Set(reflect.ValueOf(f.files[i]))
		}
		return nil
	}

	for !(i < len(f.files)) {
		if len(f.queue) == 0 {
			break
		}
		cur := f.queue[len(f.queue)-1]
		infos, err := cur.Readdir(1)
		if err != io.EOF && err != nil {
			f.err = err
			return err
		}
		for _, info := range infos {
			f.files = append(f.files, File{
				Dir:      f.queuePaths[len(f.queuePaths)-1],
				FileInfo: info,
			})
			if !info.IsDir() {
				continue
			}
			p := filepath.Join(
				f.queuePaths[len(f.queuePaths)-1],
				info.Name(),
			)
			file, err := os.Open(p)
			if err != nil {
				f.err = err
				return err
			}
			f.queue = append(f.queue, file)
			f.queuePaths = append(f.queuePaths, p)
		}
		if err == io.EOF {
			cur.Close()
			f.queue = f.queue[:len(f.queue)-1]
			f.queuePaths = f.queuePaths[:len(f.queuePaths)-1]
		}
	}

	if i >= len(f.files) {
		return ErrOutOfRange
	}
	if target != nil {
		reflect.ValueOf(target).Elem().Set(reflect.ValueOf(f.files[i]))
	}

	return nil
}

func (f *Files) Set(i int, value any) error {
	return ErrNotSupported
}
