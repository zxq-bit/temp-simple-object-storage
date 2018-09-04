package local

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"

	"github.com/caicloud/simple-object-storage/pkg/storage"
)

const (
	fileKeySplit = "-"
)

func (s *Storage) Put() (storage.Writer, error) {
	id := s.genFileID()
	key := makeFileKey(s.prefix, id)
	return NewWriter(s.base, key)
}

func (s *Storage) Get(key string) (r io.ReadCloser, e error) {
	fp := filepath.Join(s.base, key)
	f, e := os.Open(fp)
	if e != nil {
		return nil, e
	}
	return f, nil
}

func (s *Storage) Delete(key string) (e error) {
	fp := filepath.Join(s.base, key)
	return os.Remove(fp)
}

func makeFileKey(base string, id uint32) string {
	idStr := fmt.Sprintf("%08X", id)
	return strings.Join([]string{idStr, base}, fileKeySplit)
}
func (s *Storage) genFileID() uint32 {
	return atomic.AddUint32(&s.lastKey, 1)
}
