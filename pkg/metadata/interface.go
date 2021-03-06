package metadata

import "github.com/caicloud/simple-object-storage/pkg/metadata/apis"

type Bucket interface {
	ListBucket() ([]apis.Bucket, error)
	PutBucket(bucket *apis.Bucket) error
	GetBucket(name string) (*apis.Bucket, error)
	DeleteBucket(name string) error
	Close() error
}

type Object interface {
	ListObject(bucket, prefix string, start, limit int) ([]apis.Object, error)
	PutObject(object *apis.Object) error
	GetObject(bucket, key string) (*apis.Object, error)
	DeleteObject(bucket, key string) error
	Close() error
}
