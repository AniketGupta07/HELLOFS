package hellofs

import (
	"testing"

	"github.com/hanwen/go-fuse/fuse"
)

var a hello

func TestGetAttr(t *testing.T) {
	attr, err := a.GetAttr("", nil)
	if attr.Mode != fuse.S_IFDIR|0755 || err != fuse.OK {
		t.Error("Wrong")
	}
}
func TestOpenDir(t *testing.T) {
	file, err := a.OpenDir("", nil)
	if file[0].Mode != fuse.S_IFREG || err != fuse.OK {
		t.Error("Wrong")
	}
}
func (fs *hello) TestOpen(t *testing.T) {
	_, err := a.Open("hello.txt", 2, nil)
	if err != fuse.OK {
		t.Error("Wrong")
	}
}
