package slices

import "unsafe"

type eface struct {
	_    unsafe.Pointer
	data unsafe.Pointer
}

func ptrOf(obj interface{}) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(&obj)).data
}

//go:linkname noescape runtime.noescape
//go:noescape
func noescape(p unsafe.Pointer) unsafe.Pointer
