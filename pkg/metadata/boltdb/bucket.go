package boltdb

import (
	"fmt"
	"path/filepath"

	"github.com/boltdb/bolt"
	"github.com/caicloud/nirvana/log"

	"github.com/caicloud/simple-object-storage/pkg/metadata/apis"
	"github.com/caicloud/simple-object-storage/pkg/util"
)

const (
	dbFileNameBucket   = "bucket.bolt.db"
	dbBucketNameBucket = "bucket"
)

var (
	dbbbName = []byte(dbBucketNameBucket) // database bucket bucket name
)

type Bucket struct {
	db *bolt.DB
}

func NewBucket(base string, rootAllowed bool) (*Bucket, error) {
	if base = filepath.Clean(base); len(base) == 0 {
		return nil, fmt.Errorf("empty base dir")
	}
	if !rootAllowed {
		e := util.IsDeviceUnderRoot(base)
		if e != nil {
			return nil, e
		}
	}
	fp := filepath.Join(base, dbFileNameBucket)
	db, e := bolt.Open(fp, 0664, nil)
	if e != nil {
		return nil, e
	}
	b := &Bucket{
		db: db,
	}
	if e = b.init(); e != nil {
		return nil, e
	}
	return b, nil
}

func (b *Bucket) init() error {
	e := b.db.Update(func(tx *bolt.Tx) error {
		_, e := tx.CreateBucketIfNotExists(dbbbName)
		return e
	})
	return e
}

func (b *Bucket) ListBucket() ([]apis.Bucket, error) {
	// list
	var tmp []*apis.Bucket
	err := b.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket(dbbbName).ForEach(func(k, v []byte) error {
			// do it in view, in case of it may be updated
			bp, e := parseBucket(k, v)
			if e != nil {
				log.Warningf("ListBucket parseBucket %v failed, %v", string(k), e)
				return nil
			}
			tmp = append(tmp, bp)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	// pack
	re := make([]apis.Bucket, 0, len(tmp))
	for _, bp := range tmp {
		re = append(re, *bp)
	}
	return re, nil
}
func (b *Bucket) PutBucket(bp *apis.Bucket) error {
	if bp == nil {
		return fmt.Errorf("nil bucket")
	}
	k, v := packBucket(bp)
	return b.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(dbbbName).Put(k, v)
	})
}
func (b *Bucket) GetBucket(name string) (*apis.Bucket, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("empty bucket name")
	}
	var bp *apis.Bucket
	err := b.db.View(func(tx *bolt.Tx) (e error) {
		k := []byte(name)
		v := tx.Bucket(dbbbName).Get(k)
		bp, e = parseBucket(k, v)
		return e
	})
	if err != nil {
		return nil, err
	}
	return bp, nil
}
func (b *Bucket) DeleteBucket(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("empty bucket name")
	}
	return b.db.Update(func(tx *bolt.Tx) (e error) {
		return tx.Bucket(dbbbName).Delete([]byte(name))
	})
}
func (b *Bucket) Close() error {
	return b.db.Close()
}
