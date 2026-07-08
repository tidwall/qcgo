package qcgo

import (
	"unsafe"

	"github.com/tidwall/qcgo/internal"
)

// UsingFallback returns true if the slower program is using the slower fallback
// cgo interface.
func UsingFallback() bool {
	return internal.UsingFallback()
}

// CallNoStack calls function using the current Go stack, without ensuring
// addtional stack size.
func CallNoStack(cFunc, arg unsafe.Pointer) {
	internal.CallNoStack(cFunc, arg)
}

// CallSlow calls function using the standard, but slower, cgo interface.
func CallSlow(cFunc, arg unsafe.Pointer) {
	internal.CallSlow(cFunc, arg)
}

// CallWithStack calls function, ensuring the specified stack size is available.
func CallWithStack(cFunc, arg unsafe.Pointer, stackSize int) {
	if stackSize <= 0 {
		stackSize = 0
	}
	internal.CallWithStack(cFunc, arg, uintptr(stackSize))
}

// Call a function.
// Ensures at least 8K of stack size is available.
func Call(cFunc, arg unsafe.Pointer) {
	internal.Call(cFunc, arg)
}
