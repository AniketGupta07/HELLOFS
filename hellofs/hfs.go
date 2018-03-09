package hellofs

import (
	"flag"
	"log"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type hello struct {
	pathfs.FileSystem
}

func (fs *hello) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	if name == "hello.txt" {
		return &fuse.Attr{
			Mode: fuse.S_IFREG | 0644, Size: 11,
		}, fuse.OK
	}
	if name == "" {
		return &fuse.Attr{
			Mode: fuse.S_IFDIR | 0755,
		}, fuse.OK
	}
	return nil, fuse.ENOENT
}
func (fs *hello) OpenDir(name string, cont *fuse.Context) ([]fuse.DirEntry, fuse.Status) {
	if name == "" {
		p := []fuse.DirEntry{{Name: "hello.txt", Mode: fuse.S_IFREG}}
		return p, fuse.OK
	}
	return nil, fuse.ENOENT
}
func (fs *hello) Open(name string, flags uint32, cont *fuse.Context) (file nodefs.File, code fuse.Status) {
	if name != "hello.txt" {
		return nil, fuse.ENOENT
	}
	return nodefs.NewDataFile([]byte("hello world")), fuse.OK
}

//Function FS is used to create a file System
func FS() *fuse.Server {
	nfs := pathfs.NewPathNodeFs(&hello{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
	server, _, err := nodefs.MountRoot(flag.Arg(0), nfs.Root(), nil)
	if err != nil {
		log.Fatalf("ERROR MOUNTING %v", err)
	}
	return server
}
