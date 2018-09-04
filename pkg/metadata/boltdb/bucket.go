package boltdb

import "github.com/caicloud/simple-object-storage/pkg/metadata/apis"

type Bucket struct {
}

func NewBucket() (*Bucket, error) {
	return &Bucket{}, nil
}

func (b *Bucket) ListBucket() ([]apis.Bucket, error) {
	return nil, nil
}
func (b *Bucket) PutBucket(bucket *apis.Bucket) error {
	return nil
}
func (b *Bucket) GetBucket() (*apis.Bucket, error) {
	return nil, nil
}
func (b *Bucket) DeleteBucket() error {
	return nil
}
