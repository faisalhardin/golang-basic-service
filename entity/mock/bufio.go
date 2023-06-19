package mockrepo

import (
	"bufio"
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
	ScannerInterface
}

type BufioNewScanner struct {
	ScannerInterface
}

func (s *BufioNewScanner) Err() error {
	return GetBuffioErrFunc()
}

func (s *BufioNewScanner) Bytes() []byte {
	return GetBuffioBytesFunc()
}

func (s *BufioNewScanner) Text() string {
	return GetBuffioTextFunc()
}

func (s *BufioNewScanner) Scan() bool {
	return GetBuffioScanFunc()
}

func (s *BufioNewScanner) Buffer(buf []byte, max int) {
	GetBuffioBufferFunc(buf, max)
}

func (s *BufioNewScanner) Split(split bufio.SplitFunc) {
	GetBuffioSplitFunc(split)
}
