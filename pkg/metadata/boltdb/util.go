package boltdb

import (
	"encoding/json"
	"fmt"

	"github.com/caicloud/simple-object-storage/pkg/metadata/apis"
)

func bucketPrimaryKey(b *apis.Bucket) string {
	return b.Name
}
func objectPrimaryKey(o *apis.Object) string {
	return o.Key
}

func parseBucket(k, v []byte) (*apis.Bucket, error) {
	b := new(apis.Bucket) // todo object pool
	e := json.Unmarshal(v, b)
	if e != nil {
		return nil, e
	}
	if bucketPrimaryKey(b) != string(k) {
		return nil, fmt.Errorf("key value not match: %v!=%v", bucketPrimaryKey(b), string(v))
	}
	return b, nil
}
func parseObject(k, v []byte) (*apis.Object, error) {
	o := new(apis.Object) // todo object pool
	e := json.Unmarshal(v, o)
	if e != nil {
		return nil, e
	}
	if objectPrimaryKey(o) != string(k) {
		return nil, fmt.Errorf("key value not match: %v!=%v", objectPrimaryKey(o), string(v))
	}
	return o, nil
}

func packBucket(b *apis.Bucket) (k, v []byte) {
	k = []byte(bucketPrimaryKey(b))
	v, _ = json.Marshal(b)
	return
}
func packObject(o *apis.Object) (b, k, v []byte) {
	b = []byte(o.Bucket)
	k = []byte(objectPrimaryKey(o))
	v, _ = json.Marshal(o)
	return
}

func checkObjectPath(bucket, key string) error {
	if len(bucket) == 0 {
		return fmt.Errorf("empty bucket name")
	}
	if len(key) == 0 {
		return fmt.Errorf("empty object key")
	}
	return nil
}
