//go:build (arm64 || amd64) && !windows && !slowcgo

package asm

import "unsafe"

//go:nosplit
//go:noescape
func morestack()

//go:nosplit
//go:noescape
func availstack() uintptr

//go:nosplit
//go:noescape
func Call(cFunc, arg unsafe.Pointer)

//go:nosplit
func CallWithStackSize(cFunc, arg unsafe.Pointer, stacksz uintptr) {
	for availstack() < stacksz {
		morestack()
	}
	Call(cFunc, arg)
}
