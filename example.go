//go:build ignore

package main

/*
#include "qchelper.h"

int add(int a, int b) {
    return a + b;
}

QC2R(add, int, int, int)
*/
import "C"
import (
	"unsafe"

	"github.com/tidwall/qcgo"
)

func add(a, b int) int {
	args := C.add_args{arg0: C.int(a), arg1: C.int(b)}
	qcgo.Call(C.add_call, unsafe.Pointer(&args))
	return int(args.ret)
}

func main() {
	println(add(10, 20))
}

// Output: 30
