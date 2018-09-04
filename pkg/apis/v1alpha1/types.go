package v1alpha1

import "time"

type Bucket struct {
	ID   int64
	Name string
}

type Object struct {
	ID         string
	Key        string
	Bucket     string
	Size       int64
	Checksum   string
	UpdateTime time.Time
}
