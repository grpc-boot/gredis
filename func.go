package gredis

import (
	"errors"
	"unsafe"
)

func IsNil(err error) bool {
	return err == nil || errors.Is(err, ErrNil)
}

// Bytes2String converts byte slice to string.
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// String2Bytes converts string to byte slice.
func String2Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
