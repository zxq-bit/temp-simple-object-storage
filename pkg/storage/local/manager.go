package local

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/caicloud/simple-object-storage/pkg/util"
)

const (
	bucketNum = 512
)

type Storage struct {
	base string

	lastKey uint32
	prefix  string
}

func NewStorage(base string, rootAllowed bool) (*Storage, error) {
	if base = filepath.Clean(base); len(base) == 0 {
		return nil, fmt.Errorf("empty base dir")
	}
	s := &Storage{
		base: base,
	}
	if !rootAllowed {
		e := util.IsDeviceUnderRoot(s.base)
		if e != nil {
			return nil, e
		}
	}
	e := s.initDir()
	if e != nil {
		return nil, e
	}

	now := time.Now()
	rand.Seed(now.UnixNano())
	hostname, _ := os.Hostname()

	s.lastKey = rand.Uint32()
	s.prefix = fmt.Sprintf("%x", sha1.Sum([]byte(hostname+"-"+now.Format(time.RFC3339Nano))))

	return s, nil
}

func (s *Storage) initDir() error {
	for i := 0; i < bucketNum; i++ {
		dp := makeDirPath(s.base, i)
		e := os.MkdirAll(dp, 0755)
		if e != nil && !os.IsExist(e) {
			return e
		}
	}
	return nil
}

func makeDirPath(base string, id int) string {
	return filepath.Join(base, fmt.Sprintf("%08X", id))
}
