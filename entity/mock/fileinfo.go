package mockrepo

import (
	"io/fs"
	"time"
)

var GetName func() string
var GetSize func() int64
var GetMode func() fs.FileMode
var GetModTime func() time.Time
var GetIsDir func() bool
var GetSys func() interface{}

type FileInfo struct {
	fs.FileInfo
}

func (f FileInfo) Name() string {
	return GetName()
}
func (f FileInfo) Size() int64 {
	return GetSize()
}
func (f FileInfo) Mode() fs.FileMode {
	return GetMode()
}
func (f FileInfo) ModTime() time.Time {
	return GetModTime()
}
func (f FileInfo) IsDir() bool {
	return GetIsDir()
}
func (f FileInfo) Sys() interface{} {
	return GetSys()
}
