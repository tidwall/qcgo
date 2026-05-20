# qcgo

Quick CGO calls

The qcgo package provides an FFI interface for calling C functions that is 
generally faster than traditional CGO.

## When to use

For C functions that are fast. Ideally running in pure userspace.

## When not to use

For C functions that are big and slow. Like ones that do I/O, callbacks, 
or just take a long time to run.
Basically, avoid qcgo for any C code that might block other goroutines in
your Go program.

| My C function does the following        | Should I use qcgo? |
| --------------------------------------- | ------------------ |
| Accesses disk or files                  | No                 |
| Uses the network                        | No                 |
| Runs longer than a few milliseconds     | No                 |
| Has a callback to Go                    | No                 |
| Uses signals                            | No                 |
| Uses threads, pthreads, or mutexes      | No                 |
| Needs a huge stack space                | No                 |
| Accesses devices such as the GPU        | No                 |
| Does things that I'm not sure about     | No                 |

## Usage

qcgo provides the Go function `qcgo.Call(cFunc, arg)` that directly calls 
a C function using the single provided argument. 

Unless your C function has only one argument and no return value, a wrapper
function is needed. This wrapper acts as a simple trampoline for handling the 
input argument, performing the C operation, and returning a value.

The `qchelper.h` file is provided to generate wrappers.

## Example

```go
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
```

The `QC2R(add, int, int, int)` line generates the type

```C
struct add_args {
    int arg0;
    int arg1;
    int ret;
};
```

and the function

```C
void add_call(void *uargs) {
    struct add_args *args = uargs;
    args->ret = add(args->arg0, args->arg1);
}
```

Along with QC2R the following other macro generators are available.

```C
// Zero arguments, no return value
#define QC0V(name)
// Zero arguments, return value
#define QC0R(name, ret_type)
// One argument, no return value
#define QC1V(name, arg0_type)
// One argument, return value
#define QC1R(name, arg0_type, ret_type)
// Two arguments, no return value
#define QC2V(name, arg0_type, arg1_type)
// Two arguments, return value
#define QC2R(name, arg0_type, arg1_type, ret_type)
// Three arguments, no return value
#define QC3V(name, arg0_type, arg1_type, arg2_type)
// Three arguments, return value
#define QC3R(name, arg0_type, arg1_type, arg2_type, ret_type)
// Four arguments, no return value
#define QC4V(name, arg0_type, arg1_type, arg2_type, arg3_type)
// Four arguments, return value
#define QC4R(name, arg0_type, arg1_type, arg2_type, arg3_type, ret_type)
```
## Performance

Benchmark `noop` and `add` functions using Go v1.25.

```C
void noop() {
}

int add(int a, b) {
    return a + b;
}
```

- `Benchmark...CGO` uses the regular CGO calling convention
- `Benchmark...CallSlow` uses `qcgo.CallSlow()`, which falls back to regular CGO
- `Benchmark...Call` uses `qcgo.Call()`. Defaults to 8K stack
- `Benchmark...CallNoStack` using `qcgo.CallNoStack()`
- `Benchmark...GO` a pure native go function. No CGO at all. Also 'noinline'


```
go test ./... -bench .
```

#### AMD Ryzen 9 5950X (amd64)

```
goos: linux
goarch: amd64
pkg: github.com/tidwall/qcgo/tests

BenchmarkNoopCGO-32            	38697648	        30.65 ns/op
BenchmarkNoopCallSlow-32       	38437671	        30.32 ns/op
BenchmarkNoopCall-32           	239931309	         4.506 ns/op
BenchmarkNoopCallNoStack-32    	455875718	         2.397 ns/op
BenchmarkNoopGO-32             	1000000000	         1.108 ns/op

BenchmarkAddCGO-32             	36457741	        31.53 ns/op
BenchmarkAddCallSlow-32        	36290085	        32.06 ns/op
BenchmarkAddCall-32            	174842541	         6.722 ns/op
BenchmarkAddCallNoStack-32     	257389662	         4.440 ns/op
BenchmarkAddGO-32              	967067896	         1.289 ns/op
```

#### AWS c8g.8xlarge (arm64)

```
goos: linux
goarch: arm64
pkg: github.com/tidwall/qcgo/tests

BenchmarkNoopCGO-32            	18433044	        65.88 ns/op
BenchmarkNoopCallSlow-32       	18023484	        66.32 ns/op
BenchmarkNoopCall-32           	289122334	         4.170 ns/op
BenchmarkNoopCallNoStack-32    	558111374	         2.150 ns/op
BenchmarkNoopGO-32             	558057507	         2.150 ns/op

BenchmarkAddCGO-32             	16349366	        73.34 ns/op
BenchmarkAddCallSlow-32        	16783740	        71.27 ns/op
BenchmarkAddCall-32            	200903840	         5.955 ns/op
BenchmarkAddCallNoStack-32     	269759557	         4.312 ns/op
BenchmarkAddGO-32              	558058764	         2.150 ns/op
```

#### Apple M1 Max (arm64)

```
goos: darwin
goarch: arm64
pkg: github.com/tidwall/qcgo/tests

BenchmarkNoopCGO-10            	37663748	        27.91 ns/op
BenchmarkNoopCallSlow-10       	42565896	        28.18 ns/op
BenchmarkNoopCall-10           	267502512	         4.496 ns/op
BenchmarkNoopCallNoStack-10    	571079804	         2.103 ns/op
BenchmarkNoopGO-10             	1000000000	         0.9753 ns/op

BenchmarkAddCGO-10             	37455374	        32.80 ns/op
BenchmarkAddCallSlow-10        	40007277	        29.44 ns/op
BenchmarkAddCall-10            	312317916	         3.845 ns/op
BenchmarkAddCallNoStack-10     	341873401	         3.524 ns/op
BenchmarkAddGO-10              	1000000000	         0.9612 ns/op
```

## Support

Currently Linux, MacOS, and BSD running on arm64 and amd64 are supported.

All other platforms work but will fall back to using the slower CGO 
interface.

## Technical details

Under the hood, the `qcgo.Call` function runs on the current goroutine
stack. It checks that the available stack space is at least 8K, growing the 
stack as needed. It then calls a tiny asm function that in turn prepares the
arguments and calls the C function.

Additional functions include:

- `qcgo.CallWithStackSize` - Allows for larger or smaller stack sizes.
- `qcgo.CallNoStack` - Performs no stack size check at all.
- `qcgo.CallSlow` - Calls the C function using the slower traditional CGO.

## Links to similar projects

- https://words.filippo.io/rustgo
- https://github.com/petermattis/fastcgo
- https://github.com/nitrix/fastcgo
- https://github.com/ihciah/rust2go
