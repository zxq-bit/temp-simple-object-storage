package local

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/caicloud/simple-object-storage/pkg/storage"
)

const (
	fileKeySplit = "-"
)

func (s *Storage) Put() (storage.Writer, error) {
	id := s.genFileID()
	key := packFileKey(s.prefix, id)
	bucket := genBucket(id)
	return NewWriter(s.base, bucket, key)
}

func (s *Storage) Get(key string) (r io.ReadCloser, e error) {
	id, e := parseFileKey(key)
	if e != nil {
		return nil, e
	}
	bucket := genBucket(id)
	fp := filepath.Join(s.base, bucket, key)
	f, e := os.Open(fp)
	if e != nil {
		return nil, e
	}
	return f, nil
}

func (s *Storage) Delete(key string) (e error) {
	id, e := parseFileKey(key)
	if e != nil {
		return e
	}
	bucket := genBucket(id)
	fp := filepath.Join(s.base, bucket, key)
	return os.Remove(fp)
}

func packFileKey(base string, id uint32) string {
	idStr := fmt.Sprintf("%08X", id)
	return strings.Join([]string{idStr, base}, fileKeySplit)
}

func parseFileKey(key string) (uint32, error) {
	vec := strings.SplitN(key, fileKeySplit, 2)
	if len(vec) != 2 {
		return 0, fmt.Errorf("bad file key: \"%v\"", key)
	}
	id, e := strconv.Atoi(vec[0])
	if e != nil {
		return 0, fmt.Errorf("bad file key: %v", e)
	}
	return uint32(id), nil
}

func (s *Storage) genFileID() uint32 {
	return atomic.AddUint32(&s.lastKey, 1)
}

func genBucket(id uint32) string {
	return fmt.Sprintf("%04X", id%bucketNum)
}
