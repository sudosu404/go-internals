package gointernals

import (
	"unsafe"

	"github.com/sudosu404/go-internals/abi"
)

type PointerType struct {
	abi.Type
	Elem *abi.Type
}

//go:nosplit
func PointerCast[ToT any, FromT any](src *FromT) *ToT {
	return (*ToT)(unsafe.Pointer(src))
}
