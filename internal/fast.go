//go:build (arm64 || amd64) && !windows && !slowcgo

package internal

import (
	"unsafe"

	"github.com/tidwall/qcgo/internal/asm"
	"github.com/tidwall/qcgo/internal/fallback"
)

const defaultStackSize = 8 * 1024 // 8KB stack

func UsingFallback() bool {
	return false
}

func CallNoStack(cFunc, arg unsafe.Pointer) {
	asm.Call(cFunc, arg)
}

func CallSlow(cFunc, arg unsafe.Pointer) {
	fallback.Call(cFunc, arg)
}

func CallWithStack(cFunc, arg unsafe.Pointer, stackSize uintptr) {
	asm.CallWithStackSize(cFunc, arg, stackSize)
}

func Call(cFunc, arg unsafe.Pointer) {
	asm.CallWithStackSize(cFunc, arg, defaultStackSize)
}
