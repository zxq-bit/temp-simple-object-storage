package local

import (
	"crypto/md5"
	"fmt"
	"hash"
	"os"
	"path/filepath"
)

// single thread only

type Writer struct {
	base     string
	key      string
	size     int64
	checksum hash.Hash

	f *os.File
}

func NewWriter(base, bucket, key string) (*Writer, error) {
	base = filepath.Clean(base)
	bucket = filepath.Clean(bucket)
	key = filepath.Clean(key)
	fp := filepath.Join(base, bucket, key)
	f, e := os.OpenFile(fp, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0664)
	if e != nil {
		return nil, e
	}
	return &Writer{
		base:     base,
		key:      key,
		size:     0,
		checksum: md5.New(),
		f:        f,
	}, nil
}

func (w *Writer) Write(b []byte) (n int, e error) {
	if w.f == nil {
		return 0, os.ErrClosed
	}
	if n, e = w.f.Write(b); e == nil {
		_, e = w.checksum.Write(b)
	}
	w.size += int64(n)
	return n, e
}

func (w *Writer) Close() error {
	if w.f == nil {
		return nil
	}
	return w.f.Close()
}

func (w *Writer) Stat() (key string, size int64, checksum string, e error) {
	checksum = fmt.Sprintf("%x", w.checksum.Sum(nil))
	return w.key, w.size, checksum, nil
}
