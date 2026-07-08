//go:build (!arm64 && !amd64) || windows || slowcgo

package internal

import (
	"unsafe"

	"github.com/tidwall/qcgo/internal/fallback"
)

func UsingFallback() bool {
	return true
}

func CallNoStack(cFunc, arg unsafe.Pointer) {
	fallback.Call(cFunc, arg)
}

func CallSlow(cFunc, arg unsafe.Pointer) {
	fallback.Call(cFunc, arg)
}

func CallWithStack(cFunc, arg unsafe.Pointer, stackSize uintptr) {
	fallback.Call(cFunc, arg)
}

func Call(cFunc, arg unsafe.Pointer) {
	fallback.Call(cFunc, arg)
}
