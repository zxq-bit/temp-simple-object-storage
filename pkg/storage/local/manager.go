package local

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/shirou/gopsutil/disk"
)

type Storage struct {
	base      string
	bucketNum int

	lastKey uint32
	prefix  string
}

func NewStorage(base string, bucketNum int, rootAllowed bool) (*Storage, error) {
	if base = filepath.Clean(base); len(base) == 0 {
		return nil, fmt.Errorf("empty base dir")
	}
	if bucketNum < 1 {
		return nil, fmt.Errorf("bad bucket num %d", bucketNum)
	}
	s := &Storage{
		base:      base,
		bucketNum: bucketNum,
	}
	if !rootAllowed {
		e := s.checkDevice()
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
	for i := 0; i < s.bucketNum; i++ {
		dp := makeDirPath(s.base, i)
		e := os.MkdirAll(dp, 0755)
		if e != nil && !os.IsExist(e) {
			return e
		}
	}
	return nil
}

func (s *Storage) checkDevice() error {
	const root = "/"
	var (
		rootVolName string
		curVolName  string
	)
	ps, e := disk.Partitions(true)
	if e != nil {
		return e
	}
	for i := range ps {
		if ps[i].Mountpoint == root {
			rootVolName = ps[i].Device
		}
		if filepath.HasPrefix(s.base, ps[i].Mountpoint) && (len(curVolName) == 0 || ps[i].Mountpoint != root) {
			curVolName = ps[i].Device
		}
	}
	if curVolName == rootVolName {
		return fmt.Errorf("storage is under root device, \"%s\"==\"%s\"", curVolName, rootVolName)
	}
	return nil
}

func makeDirPath(base string, id int) string {
	return filepath.Join(base, fmt.Sprintf("%08X", id))
}
