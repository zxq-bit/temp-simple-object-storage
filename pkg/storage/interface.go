package storage

import (
	"io"
)

type Writer interface {
	io.WriteCloser
	Stat() (key string, size int64, checksum string, e error)
}

type Interface interface {
	Put() (w Writer, e error)
	Get(key string) (r io.ReadCloser, e error)
	Delete(key string) (e error)
}
