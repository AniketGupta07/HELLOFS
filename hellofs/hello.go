package main

import (
	"flag"
	"hellofs"
	"log"

	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type hello struct {
	pathfs.FileSystem
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatalf("MOUNT-POINT NOT SPECIFIED ")
	}
	hellofs.FS().Serve()
}
