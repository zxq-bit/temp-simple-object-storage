package apis

import "time"

type Bucket struct {
	ID    int64
	Name  string
	Owner string
}

type Object struct {
	ID         string
	Key        string
	Bucket     int64
	Size       int64
	Checksum   string
	UpdateTime time.Time

	Blocks []Block
}

type Block struct {
	Key      string
	Size     int64
	Checksum string
}
