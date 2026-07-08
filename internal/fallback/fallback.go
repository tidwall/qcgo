package fallback

/*
#include <stdint.h>
static void call(uintptr_t a, uintptr_t b) {
	((void(*)(void*))((void*)a))((void*)b);
}
*/
import "C"
import "unsafe"

func Call(cFunc, arg unsafe.Pointer) {
	C.call(C.uintptr_t(uintptr(cFunc)), C.uintptr_t(uintptr(arg)))
}
