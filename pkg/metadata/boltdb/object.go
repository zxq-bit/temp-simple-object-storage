package boltdb

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"

	"github.com/boltdb/bolt"
	"github.com/caicloud/nirvana/log"

	"github.com/caicloud/simple-object-storage/pkg/metadata/apis"
	"github.com/caicloud/simple-object-storage/pkg/util"
)

const (
	dbFileNameObject = "object.bolt.db"
)

type Object struct {
	db *bolt.DB
}

func NewObject(base string, rootAllowed bool) (*Object, error) {
	if base = filepath.Clean(base); len(base) == 0 {
		return nil, fmt.Errorf("empty base dir")
	}
	if !rootAllowed {
		e := util.IsDeviceUnderRoot(base)
		if e != nil {
			return nil, e
		}
	}
	fp := filepath.Join(base, dbFileNameObject)
	db, e := bolt.Open(fp, 0664, nil)
	if e != nil {
		return nil, e
	}
	return &Object{
		db: db,
	}, nil
}

func (o *Object) ListObject(bucket, prefix string, start, limit int) ([]apis.Object, error) {
	bucket = path.Clean(bucket)
	prefix = path.Clean(prefix)
	b := []byte(bucket)

	// list
	var tmp []*apis.Object
	err := o.db.View(func(tx *bolt.Tx) error {
		// iterate
		// prepare
		c := tx.Bucket(b).Cursor()
		iterStart := c.First
		iterCheck := func([]byte) bool { return true }
		if len(prefix) > 0 {
			pb := []byte(prefix)
			iterStart = func() (k, v []byte) { return c.Seek(pb) }
			iterCheck = func(k []byte) bool { return bytes.HasPrefix(k, pb) }
		}
		// do
		k, v := iterStart()
		// skip
		for i := 1; i < start && k != nil; i++ {
			k, v = c.Next()
		}
		// collect
		for ; k != nil && iterCheck(k) && len(tmp) < limit; c.Next() {
			op, e := parseObject(k, v)
			if e != nil {
				log.Warningf("ListObject parseObject %v/%v failed, %v", bucket, string(k), e)
				continue
			}
			tmp = append(tmp, op)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	// pack
	re := make([]apis.Object, 0, len(tmp))
	for _, op := range tmp {
		re = append(re, *op)
	}
	return re, nil
}
func (o *Object) PutObject(op *apis.Object) error {
	if op == nil {
		return fmt.Errorf("nil object")
	}
	b, k, v := packObject(op)
	return o.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(b).Put(k, v)
	})
}
func (o *Object) GetObject(bucket, key string) (*apis.Object, error) {
	if e := checkObjectPath(bucket, key); e != nil {
		return nil, e
	}
	var op *apis.Object
	err := o.db.View(func(tx *bolt.Tx) (e error) {
		b := []byte(bucket)
		k := []byte(key)
		v := tx.Bucket(b).Get(k)
		op, e = parseObject(k, v)
		return e
	})
	if err != nil {
		return nil, err
	}
	return op, nil
}
func (o *Object) DeleteObject(bucket, key string) error {
	if e := checkObjectPath(bucket, key); e != nil {
		return e
	}
	return o.db.Update(func(tx *bolt.Tx) (e error) {
		return tx.Bucket([]byte(bucket)).Delete([]byte(key))
	})
}
func (o *Object) Close() error {
	return o.db.Close()
}
