package apis

import "time"

type Bucket struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
	ACL   ACL    `json:"acl"`
}

type ACL struct {
	PublicRW uint8 `json:"publicRW"`
}

type Object struct {
	Key        string    `json:"key"`
	Bucket     string    `json:"bucket"`
	Size       int64     `json:"size"`
	Checksum   string    `json:"checksum"`
	UpdateTime time.Time `json:"updateTime"`
	Blocks     []Block   `json:"blocks"`
}

type Block struct {
	Key      string `json:"key"`
	Size     int64  `json:"size"`
	Checksum string `json:"checksum"`
}
