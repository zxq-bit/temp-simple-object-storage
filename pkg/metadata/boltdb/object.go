package boltdb

import "github.com/caicloud/simple-object-storage/pkg/metadata/apis"

type Object struct {
}

func NewObject() (*Object, error) {
	return &Object{}, nil
}

func (o *Object) ListObject(bucket, prefix string, start, limit int) ([]apis.Object, error) {
	return nil, nil
}
func (o *Object) PutObject(object *apis.Object) error {
	return nil
}
func (o *Object) GetObject(bucket, key string) (*apis.Object, error) {
	return &apis.Object{}, nil
}
func (o *Object) DeleteObject(bucket, key string) error {
	return nil
}
