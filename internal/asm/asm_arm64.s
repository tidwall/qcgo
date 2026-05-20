//go:build !windows && !slowcgo

#include "textflag.h"

TEXT ·morestack(SB), NOSPLIT|NOFRAME, $0
    JMP runtime·morestack_noctxt(SB)  // add more stack

TEXT ·availstack(SB), NOSPLIT|NOFRAME, $0-8
    MOVD    RSP, R0        // read current stack pointer
    MOVD    0(g), R1       // read stack lo boundary
    SUB     R1, R0         // avail = current - lo
    MOVD    R0, ret+0(FP)  // set avail return value
    RET                    // return

TEXT ·Call(SB), NOSPLIT|NOPTR, $0-16
    MOVD    cfunc+0(FP), R1  // cFunc unsafe.Pointer
    MOVD    arg+8(FP), R0    // arg unsafe.Pointer
    CALL    R1               // call function
    RET                      // return
