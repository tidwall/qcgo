package tests

/*
#include <stdlib.h>
#include <stdio.h>
#include <assert.h>
#include <time.h>
#include <signal.h>
#include "../qchelper.h"

int add(int a, int b) {
	return a + b;
}

void noop(void) {

}

size_t bigstack(size_t size) {
    unsigned char *data = alloca(size);
    for (size_t i = 0; i < size; i++) {
        data[i] = i;
    }
    size_t x = 0;
    for (size_t i = 0; i < size; i++) {
        x += data[i];
    }
	return x;
}

int bigsysio(const char *path, int datasize, int seconds) {
	time_t start_time = time(NULL);
	while (difftime(time(NULL), start_time) < seconds) {
		FILE *f = fopen(path, "wb+");
		assert(f);
		char *data = malloc(datasize);
		assert(data);
		for (int i = 0; i < datasize; i++) {
			data[i] = datasize+i;
		}
		assert(fwrite(data, 1, datasize, f) == datasize);
		free(data);
		fclose(f);
		raise(SIGUSR1);
	}
	remove(path);
	raise(SIGUSR2);
	return 0;
}

QC2R(add, int, int, int)
QC0V(noop)
QC1R(bigstack, size_t, size_t)
QC3R(bigsysio, const char*, int, int, int)
*/
import "C"
import (
	"os"
	"os/signal"
	"syscall"
	"unsafe"

	"github.com/tidwall/qcgo"
)

///////////////////////////////////////////////////////////////////////////////
// int add(int a, int b)
///////////////////////////////////////////////////////////////////////////////

func addCall(a, b int) int {
	args := C.add_args{C.int(a), C.int(b), 0}
	qcgo.Call(C.add_call, unsafe.Pointer(&args))
	return int(args.ret)
}

func addCallNoStack(a, b int) int {
	args := C.add_args{C.int(a), C.int(b), 0}
	qcgo.CallNoStack(C.add_call, unsafe.Pointer(&args))
	return int(args.ret)
}

func addCallSlow(a, b int) int {
	args := C.add_args{C.int(a), C.int(b), 0}
	qcgo.CallSlow(C.add_call, unsafe.Pointer(&args))
	return int(args.ret)
}

func addCGO(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

//go:noinline
func addGO(a, b int) int {
	return a + b
}

///////////////////////////////////////////////////////////////////////////////
// void noop(void)
///////////////////////////////////////////////////////////////////////////////

func noopCall() {
	qcgo.Call(C.noop_call, nil)
}

func noopCallNoStack() {
	qcgo.CallNoStack(C.noop_call, nil)
}

func noopCallSlow() {
	qcgo.CallSlow(C.noop_call, nil)
}

func noopCGO() {
	C.noop()
}

//go:noinline
func noopGO() {
}

///////////////////////////////////////////////////////////////////////////////
// size_t bigstack(size_t size)
///////////////////////////////////////////////////////////////////////////////

func bigStackCall(size int) int {
	args := C.bigstack_args{C.size_t(size), 0}
	qcgo.Call(C.bigstack_call, unsafe.Pointer(&args))
	return int(args.ret)
}

func bigStackCallNoStack(size int) int {
	args := C.bigstack_args{C.size_t(size), 0}
	qcgo.CallNoStack(C.bigstack_call, unsafe.Pointer(&args))
	return int(args.ret)
}

func bigStackCallWithStack(size int) int {
	args := C.bigstack_args{C.size_t(size), 0}
	qcgo.CallWithStack(C.bigstack_call, unsafe.Pointer(&args), size*2)
	return int(args.ret)
}

func bigStackCallSlow(size int) int {
	args := C.bigstack_args{C.size_t(size), 0}
	qcgo.CallSlow(C.bigstack_call, unsafe.Pointer(&args))
	return int(args.ret)
}

func bigStackCGO(size int) int {
	return int(C.bigstack(C.size_t(size)))
}

//go:noinline
func bigStackGO(size int) int {
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = byte(i)
	}
	var x uint64
	for i := 0; i < size; i++ {
		x += uint64(data[i])
	}
	return int(x)
}

func bigsysio(path string, datasize int, seconds int) int {
	var usr1 int
	donec := make(chan bool)
	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGUSR1, syscall.SIGUSR2)
		for s := range sigc {
			if s == syscall.SIGUSR1 {
				usr1++
			} else if s == syscall.SIGUSR2 {
				break
			}
		}
		signal.Stop(sigc)
		donec <- true
	}()
	str := C.CString(path)
	args := C.bigsysio_args{arg0: str, arg1: C.int(datasize), arg2: C.int(seconds)}
	qcgo.Call(C.bigsysio_call, unsafe.Pointer(&args))
	C.free(unsafe.Pointer(str))
	<-donec
	return usr1
}
