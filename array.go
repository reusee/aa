package aa

import "errors"

type Array interface {
	Get(index int, target any) error
	Set(index int, value any) error
}

type any = interface{}

var (
	ErrOutOfRange   = errors.New("out of range")
	ErrNotSupported = errors.New("not supported")
)
