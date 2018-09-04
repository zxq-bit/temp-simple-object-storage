package main

import (
	"log"

	"github.com/caicloud/simple-object-storage/pkg/admin"
)

func main() {
	s, e := admin.NewServer()
	if e != nil {
		log.Fatalf("NewServer failed, %v", e)
	}
	s.Run()
}
