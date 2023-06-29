// https://gitlab.com/gnufred/logslate/-/issues/99
package main

import (
	"C"
	"syscall"
	"unsafe"
)

func main() {
	var p *[2]C.int
	syscall.RawSyscall(
		syscall.SYS_PIPE2,
		uintptr(unsafe.Pointer(p)),
		uintptr(syscall.F_SETPIPE_SZ),
		1048576,
	)
}
