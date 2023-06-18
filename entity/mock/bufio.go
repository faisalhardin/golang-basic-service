package mockrepo

import (
	"bufio"
	"io"
)

var GetBuffioErrFunc func() error
var GetBuffioBytesFunc func() []byte
var GetBuffioTextFunc func() string
var GetBuffioScanFunc func() bool
var GetBuffioBufferFunc func(buf []byte, max int)
var GetBuffioSplitFunc func(split bufio.SplitFunc)

type ScannerInterface interface {
	Err() error
	Bytes() []byte
	Text() string
	Scan() bool
	Buffer(buf []byte, max int)
	Split(split bufio.SplitFunc)
}

type NewScanner interface {
	NewScanner(r io.Reader) ScannerInterface
}

type BufioNewScanner struct {
	NewScanner
}

type MockScanner struct {
	ScannerInterface
}

func (s *MockScanner) Err() error {
	return GetBuffioErrFunc()
}

func (s *MockScanner) Bytes() []byte {
	return GetBuffioBytesFunc()
}

func (s *MockScanner) Text() string {
	return GetBuffioTextFunc()
}

func (s *MockScanner) Scan() bool {
	return GetBuffioScanFunc()
}

func (s *MockScanner) Buffer(buf []byte, max int) {
	GetBuffioBufferFunc(buf, max)
}

func (s *MockScanner) Split(split bufio.SplitFunc) {
	GetBuffioSplitFunc(split)
}
